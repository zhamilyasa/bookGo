package models

import "time"

type Book struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"publishedAt"`
	Pages       int       `json:"pages"`
	Users       []User    `gorm:"many2many:user_books;"`
}
