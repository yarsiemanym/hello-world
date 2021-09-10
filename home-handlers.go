package main

import (
	"fmt"
	"net/http"
)

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /")
	fmt.Fprintf(w, "Hello, world!")
	fmt.Println("Response: Ok")
}
