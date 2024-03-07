package document

import (
	"github.com/gorilla/mux"
	"html/template"
	"huepattl.de/blog"
	"huepattl.de/blog/index"
	"huepattl.de/common"
	"huepattl.de/web/handlers"
	"log"
	"net/http"
)

type PageVariables struct {
	handlers.PageVariables
	Entry blog.Entry
	Text  template.HTML
}

func GetRequest(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	entryId := vars["entry"]
	log.Printf("Requested blog entry '%s'...", entryId)
	message := "OK"

	found, err := index.FindById(index.RepositoryProperties{
		ProjectId:  "huepattl",
		Collection: "blogs",
	}, entryId)
	if err != nil {
		message = "Failed to find blog entry"
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	documentText, err := GetDocument(found.DocumentId)
	if err != nil {
		message = "Failed to retrieve document"
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	var pageVariables = PageVariables{
		PageVariables: handlers.PageVariables{
			Title:      "Blog",
			Content:    "ignore",
			RenderedAt: "now",
		},
		Entry: *found,
	}

	htmlTemplate, err := template.ParseFiles("common/outerpage.html", "blog/document/page.html")
	if err != nil {
		message = "Failed to read templates"
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	pageVariables.Text = template.HTML(common.MarkdownToHtml(documentText))

	err = htmlTemplate.Execute(res, pageVariables)
	if err != nil {
		message = "Failed to handle request"
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	return
}
