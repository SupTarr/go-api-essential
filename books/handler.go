package books

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

		book := GetBook(db, id)
		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: book})
	}
}

func GetBooksHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		books := GetBooks(db)
		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: books})
	}
}

func CreateBooksHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var b Book
		if err := c.Bind(&b); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		CreateBook(db, &b)
		return c.JSON(http.StatusCreated, utils.Response{Status: utils.Success})
	}
}

func UpdateBooksHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		var b Book
		if err := c.Bind(&b); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		UpdateBook(db, &b)
		return c.JSON(http.StatusCreated, utils.Response{Status: utils.Success, Data: b})
	}
}

func DeleteBooksHandler(db *gorm.DB) func(c echo.Context) error {
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
