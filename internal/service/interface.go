package service

import (
	"context"

	"visagium-api/internal/domain"
)

// EmployeeService defines the service interface for Employee
type EmployeeService interface {
	RegisterEmployee(ctx context.Context, req *domain.EmployeeRequest) (*domain.EmployeeResponse, error)
	DeleteEmployee(ctx context.Context, req *domain.DeleteEmployeeRequest) (*domain.DeleteEmployeeResponse, error)
}

// AttendanceService defines the service interface for Attendance
type AttendanceService interface {
	SubmitAttendance(ctx context.Context, req *domain.AttendanceRequest) (*domain.AttendanceResponse, error)
	GetAttendanceByDate(ctx context.Context, date string) (*domain.AttendanceListResponse, error)
}
