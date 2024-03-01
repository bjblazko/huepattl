package blog

import (
	"github.com/gorilla/mux"
	"html/template"
	"huepattl.de/common"
	"huepattl.de/web/handlers"
	"log"
	"net/http"
)

type BlogEntryPageVariables struct {
	handlers.PageVariables
	Entry BlogEntry
	Text  template.HTML
}

func HandleBlogEntryGetRequest(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	entryId := vars["entry"]
	log.Printf("Requested blog entry '%s'...", entryId)
	message := "OK"

	found, err := FindById(RepositoryProperties{
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

	var pageVariables = BlogEntryPageVariables{
		PageVariables: handlers.PageVariables{
			Title:      "Blog",
			Content:    "ignore",
			RenderedAt: "now",
		},
		Entry: *found,
	}

	htmlTemplate, err := template.ParseFiles("web/templates/outerpage.html", "web/templates/blogentry.html")
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
