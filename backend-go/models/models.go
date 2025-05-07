package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"` // no exponer en JSON
	Role      string         `json:"role" gorm:"default:'owner'"`
	Orders    []Order        `json:"orders,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

// Estructura para la categoría
type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type Article struct {
	ID            uint           `json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Title         string         `json:"title" gorm:"index;not null"`
	Description   string         `json:"description"` // antes "Content"
	Price         float64        `json:"price"`
	Image         string         `json:"image"`
	Category      string         `json:"category"`
	AuthorID      *uint          `json:"authorId" gorm:"default:NULL"`
	Author        *User          `json:"author,omitempty" gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL;"`
	OrderArticles []OrderArticle `json:"orderArticles,omitempty" gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE;"`
}

type Order struct {
	ID            uint           `json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	UserID        uint           `json:"userId" gorm:"not null"`
	User          User           `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	OrderArticles []OrderArticle `json:"orderArticles,omitempty" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}

type OrderArticle struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	OrderID   uint           `json:"orderId" gorm:"not null"`
	Order     Order          `json:"-" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	ArticleID uint           `json:"articleId" gorm:"not null"`
	Article   Article        `json:"-" gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE;"`
	Quantity  int            `json:"quantity" gorm:"not null;default:1"`
}

// Migrar todas las tablas
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Article{}, &Order{}, &OrderArticle{})
}
