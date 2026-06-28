package repository

import (
	"spotsync/internal/models"

	"gorm.io/gorm"
)

type ZoneRepository interface {
	CreateZone(zone *models.ParkingZone) error
	GetAllZones() ([]models.ParkingZone, error)
	GetZoneById(id uint) (*models.ParkingZone, error)
	CountActiveReservations(zoneId uint) (int64, error)
}

type zoneRepository struct {
	db *gorm.DB
}

func NewZoneRepository(db *gorm.DB) ZoneRepository {
	return &zoneRepository{db: db}
}

func (r *zoneRepository) CreateZone(zone *models.ParkingZone) error {
	return r.db.Create(zone).Error
}

func (r *zoneRepository) GetAllZones() ([]models.ParkingZone, error) {
	var zones []models.ParkingZone
	err := r.db.Find(&zones).Error

	return zones, err
}

func (r *zoneRepository) GetZoneById(id uint) (*models.ParkingZone, error) {
	var zone models.ParkingZone
	err := r.db.First(&zone, id).Error
	if err != nil {
		return nil, err
	}
	return &zone, err
}

func (r *zoneRepository) CountActiveReservations(zoneId uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Reservation{}).Where("zone_id = ? AND status = ?", zoneId, "active").Count(&count).Error

	return count, err
}
