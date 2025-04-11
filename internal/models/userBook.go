package models

type UserBook struct {
	UserID uint `gorm:"primaryKey"`
	BookID uint `gorm:"primaryKey"`
}
