package book

import (
	"net/http"
	"strconv"

	"github.com/SupTarr/go-api-essential/utils"
	"github.com/gin-gonic/gin"
)

var books []Book = []Book{
	{ID: 1, Title: "1984", Author: "George Orwell"},
	{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: books})
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		return
	}

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: book})
			return
		}
	}

	c.JSON(http.StatusOK, utils.Response{Status: utils.DataNotFound})
}

func CreateBook(c *gin.Context) {
	book := new(Book)
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		return
	}

	book.ID = len(books) + 1
	books = append(books, *book)
	c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: book})
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		return
	}

	bookUpdate := new(Book)
	if err := c.ShouldBindJSON(&bookUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			book.Title = bookUpdate.Title
			book.Author = bookUpdate.Author
			books[i] = book
			c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: book})
			return
		}
	}

	c.JSON(http.StatusOK, utils.Response{Status: utils.DataNotFound})
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, utils.Response{Status: utils.Success})
			return
		}
	}

	c.JSON(http.StatusOK, utils.Response{Status: utils.DataNotFound})
}
