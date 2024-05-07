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
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		cookie := &http.Cookie{
			Name:    "JWT",
			Value:   t,
			Expires: time.Now().Add(3 * time.Hour),
		}

		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success})
	}
}

func ExtractUserFromJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := &UserData{}

		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*JwtCustomClaims)

		user.ID = claims.ID
		user.Email = claims.Email
		log.Printf("User id = %d, name = %s", user.ID, user.Email)

		c.Set("user", user)
		return next(c)
	}
}
