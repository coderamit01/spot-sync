package repository

import (
	"errors"
	"spotsync/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReservationRepository interface {
	CreateReservation(userId uint, zoneId uint, licensePlate string) (*models.Reservation, error)
	GetMyReservations(userId uint) ([]models.Reservation, error)
	GetReservationById(id uint) (*models.Reservation, error)
	CancelReservation(id uint) error
	GetAllReservations() ([]models.Reservation, error)
}

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) CreateReservation(userId uint, zoneId uint, licensePlate string) (*models.Reservation, error) {
	var reservation models.Reservation

	err := r.db.Transaction(func(tx *gorm.DB) error {
		var zone models.ParkingZone
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&zone, zoneId).Error; err != nil {
			return err
		}

		var count int64
		tx.Model(&models.Reservation{}).Where("zone_id = ? AND status = ?", zoneId, "active").Count(&count)

		if count >= int64(zone.TotalCapacity) {
			return errors.New("zone is full")
		}

		reservation = models.Reservation{
			UserId:       userId,
			ZoneId:       zoneId,
			LicensePlate: licensePlate,
			Status:       "active",
		}
		return tx.Create(&reservation).Error
	})
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *reservationRepository) GetMyReservations(userId uint) ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Preload("ParkingZone").Where("user_id = ?", userId).Find(&reservations).Error
	return reservations, err
}

func (r *reservationRepository) GetReservationById(id uint) (*models.Reservation, error) {
	var reservation models.Reservation
	err := r.db.First(&reservation, id).Error
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *reservationRepository) CancelReservation(id uint) error {
	return r.db.Model(&models.Reservation{}).Where("id = ?", id).Update("status", "cancelled").Error
}

func (r *reservationRepository) GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Preload("User").Preload("ParkingZone").Find(&reservations).Error
	return reservations, err
}
