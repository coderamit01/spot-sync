package service

import (
	"errors"
	"spotsync/internal/dto"
	"spotsync/internal/models"
	"spotsync/internal/repository"

	"gorm.io/gorm"
)

type ZoneService interface {
	CreateZone(req *dto.ZoneRequest) (*dto.ZoneResponse, error)
	GetAllZones() ([]dto.ZoneResponseWithAvailableSpot, error)
	GetZoneById(id uint) (*dto.ZoneResponseWithAvailableSpot, error)
}

type zoneService struct {
	zoneRepo repository.ZoneRepository
}

func NewZoneService(zoneRepo repository.ZoneRepository) ZoneService {
	return &zoneService{zoneRepo: zoneRepo}
}

func (s *zoneService) CreateZone(req *dto.ZoneRequest) (*dto.ZoneResponse, error) {

	zone := &models.ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.zoneRepo.CreateZone(zone); err != nil {
		return nil, err
	}

	return &dto.ZoneResponse{
		Id:            zone.Id,
		Name:          zone.Name,
		Type:          zone.Type,
		TotalCapacity: zone.TotalCapacity,
		PricePerHour:  zone.PricePerHour,
		CreatedAt:     zone.CreatedAt,
	}, nil

}

func (s *zoneService) GetAllZones() ([]dto.ZoneResponseWithAvailableSpot, error) {
	zones, err := s.zoneRepo.GetAllZones()
	if err != nil {
		return nil, err
	}

	var result []dto.ZoneResponseWithAvailableSpot
	for _, zone := range zones {
		count, err := s.zoneRepo.CountActiveReservations(zone.Id)
		if err != nil {
			return nil, err
		}
		result = append(result, dto.ZoneResponseWithAvailableSpot{
			Id:             zone.Id,
			Name:           zone.Name,
			Type:           zone.Type,
			TotalCapacity:  zone.TotalCapacity,
			AvailableSpots: zone.TotalCapacity - int(count),
			PricePerHour:   zone.PricePerHour,
			CreatedAt:      zone.CreatedAt,
		})
	}
	return result, nil
}

func (s *zoneService) GetZoneById(id uint) (*dto.ZoneResponseWithAvailableSpot, error) {
	zone, err := s.zoneRepo.GetZoneById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("zone not found")
		}
		return nil, err
	}

	count, err := s.zoneRepo.CountActiveReservations(zone.Id)
	if err != nil {
		return nil, err
	}

	return &dto.ZoneResponseWithAvailableSpot{
		Id:             zone.Id,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity - int(count),
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt,
	}, nil

}
