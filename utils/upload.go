package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func UploadImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: Fail, Message: err.Error()})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: Fail, Message: err.Error()})
	}
	defer src.Close()

	dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: Fail, Message: err.Error()})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: Fail, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response{Status: Success})
}
