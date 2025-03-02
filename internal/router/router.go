package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"visagium-api/internal/config"
	"visagium-api/internal/handler"
	"visagium-api/internal/repository/postgres"
	"visagium-api/internal/service"
)

// SetupRoutes initializes the routes for the API
func SetupRoutes(e *echo.Echo, db *sql.DB, cfg *config.Config) {
	// Initialize repositories
	employeeRepo := postgres.NewEmployeeRepository(db)
	attendanceRepo := postgres.NewAttendanceRepository(db)

	// Initialize services
	employeeService := service.NewEmployeeService(employeeRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo, employeeRepo)

	// Initialize handlers
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)

	// Set up routes

	// Employee routes
	employeeGroup := e.Group("/Employee")
	employeeGroup.POST("", employeeHandler.RegisterEmployee)
	employeeGroup.DELETE("", employeeHandler.DeleteEmployee)

	// Attendance routes
	attendanceGroup := e.Group("/Attendance")
	attendanceGroup.POST("", attendanceHandler.SubmitAttendance)
	attendanceGroup.GET("", attendanceHandler.GetAttendance)
}
