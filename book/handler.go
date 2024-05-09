package book

import (
	"net/http"
	"strconv"

	"github.com/SupTarr/go-api-essential/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBookHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		book, err := GetBook(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: book})
	}
}

func GetBooksHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		books, err := GetBooks(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: books})
	}
}

func CreateBookHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var b Book
		if err := c.Bind(&b); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		err := CreateBook(db, &b)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusCreated, utils.Response{Status: utils.Success})
	}
}

func UpdateBookHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		b, err := GetBook(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		if err := c.Bind(&b); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		UpdateBook(db, b)
		return c.JSON(http.StatusCreated, utils.Response{Status: utils.Success, Data: b})
	}
}

func DeleteBookHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		err = DeleteBook(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusNoContent, utils.Response{Status: utils.Success})
	}
}
