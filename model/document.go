package model

import (
	"gorm.io/gorm"
)

// Document represents a document model in the database.
type Document struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
}
