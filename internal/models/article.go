package models

import (
	"time"
)

type Article struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	Slug       string    `gorm:"unique" json:"slug"`
	Summary    string    `json:"summary"`
	Content    string    `json:"content"`
	AuthorID   uint      `json:"author_id"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Tags       []Tag     `gorm:"many2many:article_tags;" json:"tags"`
	Views      int       `json:"views"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"ID"`
	Name string `gorm:"type:varchar(255);not null;unique" json:"Name"`
	Slug string `gorm:"type:varchar(255);unique" json:"Slug"`
}

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"ID"`
	Name string `gorm:"type:varchar(255);not null;unique" json:"Name"`
	Slug string `gorm:"type:varchar(255);unique" json:"Slug"`
}
