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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	result, err := h.zoneService.CreateZone(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Parking zone created successfully",
		"data":    result,
	})

}

func (h *ZoneHandler) GetAllZones(c echo.Context) error {
	result, err := h.zoneService.GetAllZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Parking zones retrieved successfully",
		"data":    result,
	})
}

func (h *ZoneHandler) GetZoneById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid zone ID",
		})
	}

	result, err := h.zoneService.GetZoneById(uint(id))
	if err != nil {
		if err.Error() == "zone not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"success": false,
				"message": "Zone not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Parking zone retrieved successfully",
		"data":    result,
	})
}
