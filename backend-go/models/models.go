package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"uniqueIndex;not null"`
	Password string  `gorm:"not null"`
	Role     string  `gorm:"default:'owner'"`
	Orders   []Order `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type Article struct {
	gorm.Model
	Title         string  `gorm:"index;not null"`
	Content       string
	AuthorID      *uint   `gorm:"default:NULL"`
	Author        *User   `gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL;"`
	OrderArticles []OrderArticle `gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE;"`
}

type Order struct {
	gorm.Model
	UserID        uint    `gorm:"not null"`
	User          User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	OrderArticles []OrderArticle `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}

type OrderArticle struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	ArticleID uint    `gorm:"not null"`
	Article   Article `gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE;"`
	Quantity  int     `gorm:"not null;default:1"`
}

// Migrar todas las tablas
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Article{}, &Order{}, &OrderArticle{})
}