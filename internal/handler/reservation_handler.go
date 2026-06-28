package handler

import (
	"net/http"
	"spotsync/internal/dto"
	"spotsync/internal/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	reservationService service.ReservationService
}

func NewReservationHandler(reservationService service.ReservationService) *ReservationHandler {
	return &ReservationHandler{reservationService: reservationService}
}

func (h *ReservationHandler) CreateReservation(c echo.Context) error {
	userId := c.Get("userId").(uint)
	var req dto.ReservationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	result, err := h.reservationService.CreateReservation(userId, &req)
	if err != nil {
		if err.Error() == "zone is full" {
			return c.JSON(http.StatusConflict, dto.ErrorResponse{
				Success: false,
				Message: "Zone is full. No available spots.",
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Reservation confirmed successfully",
		Data:    result,
	})

}

func (h *ReservationHandler) GetMyReservations(c echo.Context) error {
	userId := c.Get("userId").(uint)

	result, err := h.reservationService.GetMyReservations(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "My reservations retrieved successfully",
		Data:    result,
	})
}

func (h *ReservationHandler) CancelReservation(c echo.Context) error {
	userId := c.Get("userId").(uint)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid reservation ID",
		})
	}

	err = h.reservationService.CancelReservation(uint(id), userId)
	if err != nil {
		if err.Error() == "forbidden" {
			return c.JSON(http.StatusForbidden, dto.ErrorResponse{
				Success: false,
				Message: "You can only cancel your own reservations",
			})
		}
		if err.Error() == "reservation not found" {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Success: false,
				Message: "Reservation not found",
			})
		}
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.ErrorResponse{
		Success: true,
		Message: "Reservation cancelled successfully",
	})

}

func (h *ReservationHandler) GetAllReservations(c echo.Context) error {
	result, err := h.reservationService.GetAllReservations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "All reservations retrieved successfully",
		Data:    result,
	})
}
