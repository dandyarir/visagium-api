package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"visagium-api/internal/domain"
	"visagium-api/internal/service"
)

// EmployeeHandler handles employee-related HTTP requests
type EmployeeHandler struct {
	employeeService service.EmployeeService
}

// NewEmployeeHandler creates a new EmployeeHandler
func NewEmployeeHandler(employeeService service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}

// RegisterEmployee handles the request to register a new employee
func (h *EmployeeHandler) RegisterEmployee(c echo.Context) error {
	req := new(domain.EmployeeRequest)
	if err := c.Bind(req); err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Invalid request format")
	}

	resp, err := h.employeeService.RegisterEmployee(c.Request().Context(), req)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteEmployee handles the request to delete an employee
func (h *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	req := new(domain.DeleteEmployeeRequest)
	if err := c.Bind(req); err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Invalid request format")
	}

	resp, err := h.employeeService.DeleteEmployee(c.Request().Context(), req)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
