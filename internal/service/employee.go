package service

import (
	"context"
	"fmt"
	"time"

	"visagium-api/internal/domain"
	"visagium-api/internal/repository"
)

// EmployeeServiceImpl is an implementation of service.EmployeeServiceImpl
type EmployeeServiceImpl struct {
	employeeRepo repository.EmployeeRepository
}

// NewEmployeeService creates a new EmployeeService
func NewEmployeeService(employeeRepo repository.EmployeeRepository) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		employeeRepo: employeeRepo,
	}
}

// RegisterEmployee registers a new employee
func (s *EmployeeServiceImpl) RegisterEmployee(ctx context.Context, req *domain.EmployeeRequest) (*domain.EmployeeResponse, error) {
	employee := &domain.Employee{
		NIP:       req.EmployeeID,
		Name:      req.Name,
		CreatedAt: time.Now(),
	}

	// Validate employee nip if exists
	emp, err := s.employeeRepo.GetByID(ctx, req.EmployeeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee by id: %w", err)
	}

	if emp != nil {
		return nil, fmt.Errorf("employee with id %d already exists", req.EmployeeID)
	}

	id, err := s.employeeRepo.Create(ctx, employee)
	if err != nil {
		return nil, fmt.Errorf("failed to register employee: %w", err)
	}

	return &domain.EmployeeResponse{
		ID:      id,
		Message: "Employee baru dibuat",
	}, nil
}

// DeleteEmployee deletes an employee
func (s *EmployeeServiceImpl) DeleteEmployee(ctx context.Context, req *domain.DeleteEmployeeRequest) (*domain.DeleteEmployeeResponse, error) {
	err := s.employeeRepo.Delete(ctx, req.EmployeeID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete employee: %w", err)
	}

	return &domain.DeleteEmployeeResponse{
		ID:      req.EmployeeID,
		Message: "employee berhasil dihapus",
	}, nil
}
