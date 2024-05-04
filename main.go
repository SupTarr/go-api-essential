package main

import (
	"net/http"

	"github.com/SupTarr/go-api-essential/book"
	"github.com/SupTarr/go-api-essential/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	r.GET("/books", book.GetBooks)
	r.GET("/books/:id", book.GetBook)
	r.POST("/books", book.CreateBook)
	r.PUT("/books/:id", book.UpdateBook)
	r.DELETE("/books/:id", book.DeleteBook)

	r.POST("/upload", uploadImage)

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	r.Run()
}

func uploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		return
	}

	err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
	}

	c.JSON(http.StatusOK, utils.Response{Status: utils.Success})
}
