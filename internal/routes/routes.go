package routes

import (
	"net/http"
	"spotsync/internal/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, authHandler *handler.AuthHandler) {

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "SpotSync API is running",
		})
	})

	api := e.Group("/api/v1")

	auth := api.Group("/auth")

	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
}
