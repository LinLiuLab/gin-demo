package main

import (
	"gin-demo/handler"
	"gin-demo/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeDatabase() {
	connection_string := "root:root@tcp(localhost:3306)/gindemo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(connection_string), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Document{})
}

func main() {
	InitializeDatabase()
	r := gin.Default()
	r.GET("/document/:id", func(c *gin.Context) {
		handler.GetDocumentHandler(c, db)
	})
	r.GET("/document", func(c *gin.Context) {
		handler.GetAllDocumentsHandler(c, db)
	})
	r.POST("/document", func(c *gin.Context) {
		handler.CreateDocumentHandler(c, db)
	})
	r.PUT("/document/:id", func(c *gin.Context) {
		handler.UpdateDocumentHandler(c, db)
	})
	r.DELETE("/document/:id", func(c *gin.Context) {
		handler.DeleteDocumentHandler(c, db)
	})
	r.Run(":8000") // listen and serve on localhost:8000
}
