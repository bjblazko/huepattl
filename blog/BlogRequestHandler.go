package blog

import (
	"html/template"
	"huepattl.de/web/handlers"
	"log"
	"net/http"
)

type BlogPageVariables struct {
	handlers.PageVariables
	Blogs []Entry
}

func HandleBlogGetRequest(res http.ResponseWriter, req *http.Request) {

	list()
	htmlTemplate, err := template.ParseFiles("web/templates/outerpage.html", "web/templates/bloglist.html")
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal("Failed to read templates", err)
		return
	}

	pageVariables := BlogPageVariables{
		PageVariables: handlers.PageVariables{
			Title:      "Blog",
			Content:    "ignore",
			RenderedAt: "now",
		},
		Blogs: list(),
	}

	err = htmlTemplate.Execute(res, pageVariables)

	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal("Failed to handle response", err)
		return
	}

	return
}

func list() []Entry {
	datasource := BigQueryTable{ProjectId: "huepattl", Dataset: "site", Table: "blogs"}
	connection, err := CreateClient(datasource.ProjectId)
	if err != nil {
		log.Fatalf("Failed to connect to BigQuery %w", err)
	}

	blogEntries, err := List(connection, datasource)
	if err != nil {
		log.Fatalf("Failed find blog listings %w", err)
	}

	return blogEntries
}
