package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", GetHomeHandler).Methods("GET")
	myRouter.HandleFunc("/articles", GetArticlesHandler).Methods("GET")
	myRouter.HandleFunc("/articles/{slug}", GetArticleHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	handleRequests()
}
