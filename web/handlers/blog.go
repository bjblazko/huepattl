package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Blog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello blog")
}

func BlogByDate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprintf(res, "Hello blog %s", vars["date"])
}