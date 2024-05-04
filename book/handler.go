package book

import (
	"net/http"
	"strconv"

	"github.com/SupTarr/go-api-essential/utils"
	"github.com/labstack/echo/v4"
)

var books []Book = []Book{
	{ID: 1, Title: "1984", Author: "George Orwell"},
	{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
}

func GetBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: books})
}

func GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
	}

	for _, book := range books {
		if book.ID == id {
			return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: book})
		}
	}

	return c.JSON(http.StatusOK, utils.Response{Status: utils.DataNotFound})
}

func CreateBook(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
	}

	book.ID = len(books) + 1
	books = append(books, *book)
	return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: book})
}

func UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
	}

	bookUpdate := new(Book)
	if err := c.Bind(&bookUpdate); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
	}

	for i, book := range books {
		if book.ID == id {
			book.Title = bookUpdate.Title
			book.Author = bookUpdate.Author
			books[i] = book
			return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: book})
		}
	}

	return c.JSON(http.StatusOK, utils.Response{Status: utils.DataNotFound})
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.JSON(http.StatusOK, utils.Response{Status: utils.Success})
		}
	}

	return c.JSON(http.StatusOK, utils.Response{Status: utils.DataNotFound})
}
