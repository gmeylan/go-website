package markdown

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gmeylan/go-website/internal/types"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func ParseBlogPost(filePath string) (types.BlogPost, error) {
	content, err := os.ReadFile(filePath)
	metadata, mainContent, err := extractFrontMatter(content)
	fmt.Println("Metadata:", metadata)
	fmt.Println("mainContent:", mainContent)

	if err != nil {
		return types.BlogPost{}, fmt.Errorf("error while reading file: %w", err)
	}

	var post types.BlogPost

	htmlContent := MarkdownToHTML(mainContent)
	post.Content = htmlContent

	if title, ok := metadata["title"]; ok {
		post.Title = strings.TrimSpace(title)
	}

	if author, ok := metadata["author"]; ok {
		post.Author = strings.TrimSpace(author)
	}

	if date, ok := metadata["published-on"]; ok {
		// Essayer plusieurs formats de date
		formats := []string{
			"2 January 2006",
			"January 2, 2006",
			"2006-01-02",
			"02/01/2006",
			"01/02/2006", // Format US
			"2006/01/02",
		}

		var parseErr error
		for _, format := range formats {
			parsedDate, err := time.Parse(format, strings.TrimSpace(date))
			if err == nil {
				post.Date = parsedDate
				parseErr = nil
				break
			}
			parseErr = err
		}

		if parseErr != nil {
			post.Date = time.Now()
		}
	} else {
		post.Date = time.Now()
	}

	return post, nil
}

func MarkdownToHTML(md []byte) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs

	p := parser.NewWithExtensions(extensions)

	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

func extractFrontMatter(content []byte) (map[string]string, []byte, error) {
	text := string(content)

	if !strings.HasPrefix(text, "---\n") {
		return make(map[string]string), content, nil
	}

	endPos := strings.Index(text[4:], "---\n")
	if endPos == -1 {
		return make(map[string]string), content, fmt.Errorf("délimiteur de fin de frontmatter non trouvé")
	}

	endPos += 4

	frontMatter := text[4:endPos]

	mainContent := text[endPos+4:] // +4 pour sauter le délimiteur de fin

	metadata := parseFrontMatter(frontMatter)

	return metadata, []byte(mainContent), nil
}

func parseFrontMatter(frontMatter string) map[string]string {
	metadata := make(map[string]string)

	lines := strings.Split(frontMatter, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		colonPos := strings.Index(line, ":")
		if colonPos == -1 {
			continue
		}

		key := strings.TrimSpace(line[:colonPos])
		value := strings.TrimSpace(line[colonPos+1:])

		metadata[key] = value
	}

	return metadata
}
