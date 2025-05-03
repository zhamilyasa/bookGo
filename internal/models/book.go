package models

import "time"

type Book struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Pages       int       `json:"pages"`
	PublishedAt time.Time `json:"publishedAt"`
	CreatorID   uint      `json:"creatorId"` //
	Users       []User    `gorm:"many2many:user_books;"`
}
