package index

import (
	"html/template"
	"huepattl.de/blog"
	"huepattl.de/web/handlers"
	"log"
	"net/http"
)

type PageVariables struct {
	handlers.PageVariables
	Blogs []blog.Entry
}

func GetRequest(res http.ResponseWriter, req *http.Request) {
	log.Printf("Requested blog listing...")
	message := "OK"

	blogs, err := list()
	if err != nil {
		message = "Failed to read blog list"
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	htmlTemplate, err := template.ParseFiles("common/outerpage.html", "blog/index/page.html")
	if err != nil {
		message = "Failed to read templates"
		log.Print(message, err)
		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	pageVariables := PageVariables{
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

func list() ([]blog.Entry, error) {
	datasource := RepositoryProperties{ProjectId: "huepattl", Collection: "blogs"}

	return List(datasource)
}
