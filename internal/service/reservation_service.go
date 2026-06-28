package service

import (
    "errors"
    "spotsync/internal/dto"
    "spotsync/internal/repository"
)

type ReservationService interface {
    CreateReservation(userId uint, req *dto.ReservationRequest) (*dto.ReservationResponse, error)
    GetMyReservations(userId uint) ([]dto.MyReservationResponse, error)
    CancelReservation(reservationId uint, userId uint) error
    GetAllReservations() ([]dto.MyReservationResponse, error)
}

type reservationService struct {
    reservationRepo repository.ReservationRepository
}

func NewReservationService(reservationRepo repository.ReservationRepository) ReservationService {
    return &reservationService{reservationRepo: reservationRepo}
}

func (s *reservationService) CreateReservation(userId uint, req *dto.ReservationRequest) (*dto.ReservationResponse, error) {
    reservation, err := s.reservationRepo.CreateReservation(userId, req.ZoneId, req.LicensePlate)
    if err != nil {
        return nil, err
    }

    return &dto.ReservationResponse{
        Id:           reservation.Id,
        UserId:       reservation.UserId,
        ZoneId:       reservation.ZoneId,
        LicensePlate: reservation.LicensePlate,
        Status:       reservation.Status,
        CreatedAt:    reservation.CreatedAt,
        UpdatedAt:    reservation.UpdatedAt,
    }, nil
}

func (s *reservationService) GetMyReservations(userId uint) ([]dto.MyReservationResponse, error) {
    reservations, err := s.reservationRepo.GetMyReservations(userId)
    if err != nil {
        return nil, err
    }

    var result []dto.MyReservationResponse
    for _, r := range reservations {
        result = append(result, dto.MyReservationResponse{
            Id:           r.Id,
            LicensePlate: r.LicensePlate,
            Status:       r.Status,
            Zone: dto.ZoneInfo{
                Id:   r.ParkingZone.Id,
                Name: r.ParkingZone.Name,
                Type: r.ParkingZone.Type,
            },
            CreatedAt: r.CreatedAt,
        })
    }
    return result, nil
}

func (s *reservationService) CancelReservation(reservationId uint, userId uint) error {
    reservation, err := s.reservationRepo.GetReservationById(reservationId)
    if err != nil {
        return errors.New("reservation not found")
    }

    // Only owner can cancel
    if reservation.UserId != userId {
        return errors.New("forbidden")
    }

    // Only active reservations can be cancelled
    if reservation.Status != "active" {
        return errors.New("only active reservations can be cancelled")
    }

    return s.reservationRepo.CancelReservation(reservationId)
}

func (s *reservationService) GetAllReservations() ([]dto.MyReservationResponse, error) {
    reservations, err := s.reservationRepo.GetAllReservations()
    if err != nil {
        return nil, err
    }

    var result []dto.MyReservationResponse
    for _, r := range reservations {
        result = append(result, dto.MyReservationResponse{
            Id:           r.Id,
            LicensePlate: r.LicensePlate,
            Status:       r.Status,
            Zone: dto.ZoneInfo{
                Id:   r.ParkingZone.Id,
                Name: r.ParkingZone.Name,
                Type: r.ParkingZone.Type,
            },
            CreatedAt: r.CreatedAt,
        })
    }
    return result, nil
}
