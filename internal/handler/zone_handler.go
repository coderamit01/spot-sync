package handler

import (
	"net/http"
	"spotsync/internal/dto"
	"spotsync/internal/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ZoneHandler struct {
	zoneService service.ZoneService
}

func NewZoneHandler(zoneService service.ZoneService) *ZoneHandler {
	return &ZoneHandler{zoneService: zoneService}
}

func (h *ZoneHandler) CreateZone(c echo.Context) error {
	var req dto.ZoneRequest

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

	result, err := h.zoneService.CreateZone(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Parking zone created successfully",
		Data:    result,
	})

}

func (h *ZoneHandler) GetAllZones(c echo.Context) error {
	result, err := h.zoneService.GetAllZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Parking zones retrieved successfully",
		Data:    result,
	})
}

func (h *ZoneHandler) GetZoneById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid zone ID",
		})
	}

	result, err := h.zoneService.GetZoneById(uint(id))
	if err != nil {
		if err.Error() == "zone not found" {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Success: false,
				Message: "Zone not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Parking zone retrieved successfully",
		Data:    result,
	})
}
