package markdown

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
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
		formats := []string{
			"2 January 2006",
			"January 2, 2006",
			"2006-01-02",
			"02/01/2006",
			"01/02/2006",
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

	if tags, ok := metadata["tags"]; ok {
		post.Tags = strings.Split(tags, ",")
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
		return make(map[string]string), content, fmt.Errorf("no end to frontmatter found")
	}

	endPos += 4

	frontMatter := text[4:endPos]

	mainContent := text[endPos+4:]

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

func GetAllBlogPosts(postsDir string, logger *slog.Logger) ([]types.BlogPost, error) {
	logger.Info(postsDir)

	var posts []types.BlogPost

	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération du répertoire de travail: %w", err)
	}

	fullPath := filepath.Join(workDir, postsDir)

	err = filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error("Error navigating in the direction", "error", err, "path", path)
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			logger.Info("Parsing Mardown file", "file", path)
			post, err := ParseBlogPost(path)
			if err != nil {
				logger.Error("Error while parsing file", "error", err, "file", path)
				return err
			}
			posts = append(posts, post)
		}

		return nil
	})

	return posts, nil
}

func GetAllTags(posts []types.BlogPost) []types.TagInfo {
	tagMap := make(map[string]*types.TagInfo)

	for _, post := range posts {
		for _, tag := range post.Tags {
			cleanTag := strings.TrimSpace(tag)
			if cleanTag == "" {
				continue
			}

			if _, exists := tagMap[cleanTag]; !exists {
				tagMap[cleanTag] = &types.TagInfo{
					Name:  cleanTag,
					Count: 0,
					Posts: []types.BlogPost{},
				}
			}

			tagMap[cleanTag].Count++
			tagMap[cleanTag].Posts = append(tagMap[cleanTag].Posts, post)
		}
	}

	var tags []types.TagInfo
	for _, tagInfo := range tagMap {
		tags = append(tags, *tagInfo)
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	})

	return tags
}
