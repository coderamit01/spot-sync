package routes

import (
	"spotsync/internal/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, authHandler *handler.AuthHandler) {
	api := e.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
}
