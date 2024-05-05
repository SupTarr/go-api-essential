package user

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SupTarr/go-api-essential/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request LoginRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		if request.Email != user.Email || request.Password != user.Password {
			return c.JSON(http.StatusUnauthorized, utils.Response{Status: utils.Fail})
		}

		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			log.Fatal("SECRET_KEY is not set")
		}

		claims := &JwtCustomClaims{
			user.Email,
			"admin",
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success, Data: map[string]string{"token": t}})
	}
}

func ExtractUserFromJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := &UserData{}

		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*JwtCustomClaims)

		user.Email = claims.Email
		user.Role = claims.Role
		log.Printf("User name = %s, role = %s", user.Email, user.Role)

		c.Set("user", user)
		return next(c)
	}
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*UserData)
		if user.Role != "admin" {
			return c.JSON(http.StatusUnauthorized, utils.Response{Status: utils.Fail})
		}

		return next(c)
	}
}
