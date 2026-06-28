package routes

import (
	"net/http"
	"spotsync/internal/handler"
	"spotsync/internal/middleware"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, authHandler *handler.AuthHandler, zoneHandler *handler.ZoneHandler) {

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "SpotSync API is running",
		})
	})

	api := e.Group("/api/v1")

	//Auth
	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	//Zones
	zones := api.Group("/zones")
	zones.GET("", zoneHandler.GetAllZones)
	zones.GET("/:id", zoneHandler.GetZoneById)
	zones.POST("", zoneHandler.CreateZone, middleware.AuthMiddleware, middleware.AdminOnly)
}
