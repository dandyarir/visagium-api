package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"visagium-api/internal/domain"
	"visagium-api/internal/service"
)

// AttendanceHandler handles attendance-related HTTP requests
type AttendanceHandler struct {
	attendanceService service.AttendanceService
}

// NewAttendanceHandler creates a new AttendanceHandler
func NewAttendanceHandler(attendanceService service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
	}
}

// SubmitAttendance handles the request to submit attendance
func (h *AttendanceHandler) SubmitAttendance(c echo.Context) error {
	req := new(domain.AttendanceRequest)
	if err := c.Bind(req); err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Invalid request format")
	}

	resp, err := h.attendanceService.SubmitAttendance(c.Request().Context(), req)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// GetAttendance handles the request to get attendance data by date
func (h *AttendanceHandler) GetAttendance(c echo.Context) error {
	date := c.QueryParam("date")
	if date == "" {
		return RespondWithError(c, http.StatusBadRequest, "Date parameter is required")
	}

	resp, err := h.attendanceService.GetAttendanceByDate(c.Request().Context(), date)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
