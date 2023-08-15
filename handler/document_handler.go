package handler

import (
	"gin-demo/database"
	"gin-demo/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDocumentHandler(c *gin.Context, db *gorm.DB) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	published := c.PostForm("published")
	document := model.Document{
		Title:     title,
		Content:   content,
		Published: published == "true",
	}
	err := database.CreateDocument(&document, db)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating document: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Document created successfully.",
	})
}

func GetDocumentHandler(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	documentId, err := strconv.Atoi(id)
	document, err := database.GetDocument(documentId, db)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error retrieving document: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Document retrieved successfully.",
		"document": document,
	})
}

func GetAllDocumentsHandler(c *gin.Context, db *gorm.DB) {
	documents, err := database.GetAllDocuments(db)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error retrieving documents: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":   "Documents retrieved successfully.",
		"documents": documents,
	})
}

func UpdateDocumentHandler(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	documentId, err := strconv.Atoi(id)
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(500, gin.H{
			"message": "Error binding json: " + err.Error(),
		})
		return
	}
	title := jsonData["title"].(string)
	content := jsonData["content"].(string)
	published := jsonData["published"].(bool)
	document := model.Document{
		Title:     title,
		Content:   content,
		Published: published,
	}
	err = database.UpdateDocument(documentId, &document, db)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error updating document: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Document updated successfully.",
	})
}

func DeleteDocumentHandler(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	documentId, err := strconv.Atoi(id)
	err = database.DeleteDocument(documentId, db)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error deleting document: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Document deleted successfully.",
	})
}
