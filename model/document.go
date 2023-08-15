package model

import (
	"gorm.io/gorm"
)

// Document represents a document model in the database.
type Document struct {
	gorm.Model
	Title     string
	Content   string
	Published bool
}
