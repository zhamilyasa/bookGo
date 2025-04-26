package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"default:user" json:"role"` // "user" или "admin"
	Books    []Book `gorm:"many2many:user_books;" json:"-"`
}
