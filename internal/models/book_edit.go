package models

type BookEdit struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"publishedAt"`
	Pages       int    `json:"pages"`
}
