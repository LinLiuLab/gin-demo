package database

import (
	"fmt"
	"gin-demo/model"
	"gorm.io/gorm"
)

func CreateDocument(document *model.Document, db *gorm.DB) error {
	fmt.Println("CreateDocument {}", document.Title)
	return db.Create(document).Error
}

func GetDocument(id int, db *gorm.DB) (*model.Document, error) {
	var document model.Document
	err := db.First(&document, id).Error
	return &document, err
}

func GetAllDocuments(db *gorm.DB) (*[]model.Document, error) {
	var documents []model.Document
	err := db.Find(&documents).Error
	return &documents, err
}

func UpdateDocument(id int, document *model.Document, db *gorm.DB) error {
	return db.Model(&model.Document{}).Where("id = ?", id).Updates(document).Error
}

func DeleteDocument(id int, db *gorm.DB) error {
	return db.Delete(&model.Document{}, id).Error
}
