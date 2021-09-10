package main

type Article struct {
	Slug   string `json:"slug"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}
