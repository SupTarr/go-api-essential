package user

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SupTarr/go-api-essential/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		err := CreateUser(db, user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success})
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request LoginRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		user, err := GetUser(db, request.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, utils.Response{Status: utils.Fail})
		}

		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			log.Fatal("SECRET_KEY is not set")
		}

		claims := &JwtCustomClaims{
			user.ID,
			user.Email,
			jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    t,
			Expires:  time.Now().Add(3 * time.Hour),
			HttpOnly: true,
		}

		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success})
	}
}

func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			log.Fatal("SECRET_KEY is not set")
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return next(c)
	}
}
