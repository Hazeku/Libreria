package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

type Article struct {
	ID      uint   `gorm:"primaryKey"`
	Title   string
	Content string
}