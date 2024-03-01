package blog

import (
	"html/template"
	"huepattl.de/web/handlers"
	"log"
	"net/http"
)

type BlogListPageVariables struct {
	handlers.PageVariables
	Blogs []BlogEntry
}

func HandleBlogListGetRequest(res http.ResponseWriter, req *http.Request) {
	log.Printf("Requested blog listing...")
	message := "OK"

	blogs, err := list()
	if err != nil {
		message = "Failed to read blog list"
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	htmlTemplate, err := template.ParseFiles("web/templates/outerpage.html", "web/templates/bloglist.html")
	if err != nil {
		message = "Failed to read templates"
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	pageVariables := BlogListPageVariables{
		PageVariables: handlers.PageVariables{
			Title:      "Blog",
			Content:    "ignore",
			RenderedAt: "now",
		},
		Blogs: blogs,
	}

	err = htmlTemplate.Execute(res, pageVariables)

	if err != nil {
		message = err.Error()
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	return
}

func list() ([]BlogEntry, error) {
	datasource := RepositoryProperties{ProjectId: "huepattl", Collection: "blogs"}

	return List(datasource)
}
