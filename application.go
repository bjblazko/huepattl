package main

import (
	"fmt"
	"github.com/gorilla/mux"
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
	router.HandleFunc("/", handlers.Show).Methods("GET")
	router.HandleFunc("/doc/{name}", handlers.Show).Methods("GET")
	router.HandleFunc("/blog", handlers.Blog).Methods("GET")
	router.HandleFunc("/blog/{date}", handlers.BlogByDate).Methods("GET")
}
