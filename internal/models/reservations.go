package models

import "time"

type Reservation struct {
	Id           uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId       uint        `gorm:"not null" json:"user_id"`
	ZoneId       uint        `gorm:"not null" json:"zone_id"`
	LicensePlate string      `gorm:"not null;size:15" json:"license_plate"`
	Status       string      `gorm:"not null;default:active" json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	User         User        `gorm:"foreignKey:UserId" json:"user,omitempty"`
	ParkingZone  ParkingZone `gorm:"foreignKey:ZoneId" json:"zone,omitempty"`
}
