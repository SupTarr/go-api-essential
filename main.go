package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/SupTarr/go-api-essential/book"
	"github.com/SupTarr/go-api-essential/user"
	"github.com/SupTarr/go-api-essential/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
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

	e.POST("/login", user.Login())
	e.POST("/upload", utils.UploadImage)
	e.GET("/index", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "World")
	})

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not set")
	}

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(user.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
	}

	bookGroup := e.Group("/books")
	bookGroup.Use(user.IsAdmin)
	bookGroup.Use(echojwt.WithConfig(config))
	bookGroup.Use(user.ExtractUserFromJWT)
	bookGroup.GET("/", book.GetBooks)
	bookGroup.GET("/:id", book.GetBook)
	bookGroup.POST("/", book.CreateBook)
	bookGroup.PUT("/:id", book.UpdateBook)
	bookGroup.DELETE("/:id", book.DeleteBook)

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
