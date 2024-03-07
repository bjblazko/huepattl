package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"huepattl.de/blog/document"
	"huepattl.de/blog/index"
	"huepattl.de/web/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	router := mux.NewRouter()

	setupRoutes(router)

	http.Handle("/", router)

	log.Fatal(
		http.ListenAndServe(":8080", nil))
}

func setupRoutes(router *mux.Router) {
	staticDir := "/static/"
	router.PathPrefix("/static/").Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("web/static"))))

	// Blog
	router.HandleFunc("/blog", index.GetRequest).Methods("GET")
	router.HandleFunc("/blog/{entry}", document.GetRequest).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/blog", http.StatusFound)
	}).Methods("GET")

	router.HandleFunc("/imprint", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/doc/imprint", http.StatusFound)
	}).Methods("GET")
	router.HandleFunc("/doc/{name}", handlers.Document).Methods("GET")
}
