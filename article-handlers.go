package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var articles = map[string]Article{
	"test": {
		Slug:   "test",
		Title:  "Golang REST API – Getting Started",
		Author: "GoDocs",
		Link:   "https://golangdocs.com/golang-rest-api",
	},
}

func GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /articles")
	values := make([]Article, 0)

	for _, value := range articles {
		values = append(values, value)
	}

	json.NewEncoder(w).Encode(values)
	fmt.Println("Response: Ok")
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	fmt.Println("Endpoint Hit: /articles/" + slug)

	article, exists := articles[slug]

	if !exists {
		w.WriteHeader(404)
		fmt.Println("Response: NotFound")
		return
	}

	json.NewEncoder(w).Encode(article)
	fmt.Println("Response: Ok")
}
