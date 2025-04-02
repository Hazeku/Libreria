package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"default:'owner'"` // Puedes cambiarlo a 'admin' si necesitas más roles
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Article struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"index;not null"` // Indexado para mejor búsqueda
	Content   string    `gorm:"not null"`
	AuthorID  uint      `gorm:"not null"` // Relacionado con User
	Author    User      `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
