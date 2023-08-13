package api_route_private

import (
	"golang-practive-webapi/src/model"

	"github.com/labstack/echo/v4"
)

func V1GetIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")

		id, err := model.ParseMessageID(paramID);
		if  err != nil {
			c.Logger().Warn(err)
			return echo.ErrNotFound
		}
		message := &model.Message{
			ID: id,
			Content: "This is PRIVATE message.",
		}

		return c.JSON(200, message)
	}
}
