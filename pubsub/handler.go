package pubsub

import (
	"net/http"

	"github.com/SupTarr/go-api-essential/utils"
	"github.com/labstack/echo/v4"
)

func PublishHandler(ps *PubSub) echo.HandlerFunc {
	return func(c echo.Context) error {
		msg := new(Message)
		if err := c.Bind(&msg); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: utils.Fail, Message: err.Error()})
		}

		go ps.Publish(msg)
		return c.JSON(http.StatusOK, utils.Response{Status: utils.Success})
	}
}
