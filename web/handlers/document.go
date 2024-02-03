package handlers

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type PageVariables struct {
	Title      string
	Content    template.HTML
	RenderedAt string
}

func Document(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	log.Printf("Requested document '%s'", name)

	staticPages := map[string]string{
		"404":     "web/static/markdown/404.md",
		"index":   "web/static/markdown/index.md",
		"test":    "web/static/markdown/test.md",
		"imprint": "web/static/markdown/imprint.md",
		"blog":    "web/static/markdown/blog.md",
	}

	htmlTemplate, err := template.ParseFiles("web/templates/outerpage.html", "web/templates/document.html")
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal("Failed to read templates", err)
		return
	}

	staticFile, found := staticPages[name]
	if !found {
		staticFile = staticPages["404"]
	}

	testMarkdown, err := readFileToString(staticFile)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal("Failed to read test file", err)
		return
	}

	variables := PageVariables{
		Title:      fmt.Sprintf("huepattl.de - %s", "header"),
		Content:    template.HTML(markdownToHtml(testMarkdown)),
		RenderedAt: getNow(),
	}

	err = htmlTemplate.Execute(res, variables)

	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal("Failed to handle response", err)
		return
	}

}

func getNow() string {
	return time.Now().Format(time.RFC1123)
}

func markdownToHtml(md string) string {
	markdownBytes := []byte(md)

	p := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock)
	doc := p.Parse(markdownBytes)
	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags | html.HrefTargetBlank})
	htmlbytes := markdown.Render(doc, renderer)
	return string(htmlbytes)
}

func readFileToString(filePath string) (string, error) {
	// Read the entire file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	fileContent := string(content)

	return fileContent, nil
}
