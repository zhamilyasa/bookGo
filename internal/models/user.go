package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Books    []Book `gorm:"many2many:user_books;"`
}
