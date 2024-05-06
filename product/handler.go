package product

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/SupTarr/go-api-essential/utils"
	"github.com/labstack/echo/v4"
)

func GetProductHandler(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		product, err := GetProduct(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: product})
	}
}

func GetProductsHandler(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		products, err := GetProducts(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: products})
	}
}

func CreateProductsHandler(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var p Product
		if err := c.Bind(&p); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		err := CreateProduct(db, &p)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusCreated, utils.Response{Status: utils.Success})
	}
}

func UpdateProductsHandler(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		var p Product
		if err := c.Bind(&p); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		product, err := UpdateProduct(db, id, &p)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusCreated, utils.Response{Status: utils.Success, Data: product})
	}
}

func DeleteProductsHandler(db *sql.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		err = DeleteProduct(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusNoContent, utils.Response{Status: utils.Success})
	}
}
