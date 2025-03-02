package repository

import (
	"context"

	"visagium-api/internal/domain"
)

// EmployeeRepository defines the repository interface for Employee
type EmployeeRepository interface {
	Create(ctx context.Context, employee *domain.Employee) (int64, error)
	Delete(ctx context.Context, employeeID int64) error
	GetByID(ctx context.Context, employeeID int64) (*domain.Employee, error)
}

// AttendanceRepository defines the repository interface for AttendanceLog
type AttendanceRepository interface {
	Create(ctx context.Context, attendance *domain.AttendanceLog) error
	GetByDate(ctx context.Context, date string) ([]domain.AttendanceData, error)
}
