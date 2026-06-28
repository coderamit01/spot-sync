package main

import (
	"log"
	"os"
	"spotsync/internal/config"
	"spotsync/internal/handler"
	"spotsync/internal/repository"
	"spotsync/internal/routes"
	"spotsync/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	godotenv.Load()
	config.ConnectDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}
	//Auth
	userRepo := repository.NewUserRepository(config.DB)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	//Zone
	zoneRepo := repository.NewZoneRepository(config.DB)
	zoneService := service.NewZoneService(zoneRepo)
	zoneHandler := handler.NewZoneHandler(zoneService)
	//Reservation
	reservationRepo := repository.NewReservationRepository(config.DB)
	reservationService := service.NewReservationService(reservationRepo)
	reservationHandler := handler.NewReservationHandler(reservationService)

	routes.Routes(e, authHandler, zoneHandler, reservationHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(e.Start(":" + port))
}
