package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/SupTarr/go-api-essential/book"
	"github.com/SupTarr/go-api-essential/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	t := &utils.Template{
		Templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	e.Renderer = t

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	e.GET("/books", book.GetBooks)
	e.GET("/books/:id", book.GetBook)
	e.POST("/books", book.CreateBook)
	e.PUT("/books/:id", book.UpdateBook)
	e.DELETE("/books/:id", book.DeleteBook)

	e.POST("/upload", utils.UploadImage)
	e.GET("/index", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "World")
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
