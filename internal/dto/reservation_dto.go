package dto

import "time"

type ReservationRequest struct {
	ZoneId       uint   `json:"zone_id"       validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required,max=15"`
}

type ReservationResponse struct {
	Id           uint      `json:"id"`
	UserId       uint      `json:"user_id"`
	ZoneId       uint      `json:"zone_id"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ZoneInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type MyReservationResponse struct {
	Id           uint      `json:"id"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	Zone         ZoneInfo  `json:"zone"`
	CreatedAt    time.Time `json:"created_at"`
}
