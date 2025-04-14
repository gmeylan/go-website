package markdown

import (
	"fmt"
	"os"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type BlogPost struct {
	Title       string
	Date        time.Time
	Description string
	Content     string
	Slug        string
}

func ParseMarkdownFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error while reading file: %w", err)
	}

	htmlContent := MarkdownToHTML(content)
	return htmlContent, nil
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
