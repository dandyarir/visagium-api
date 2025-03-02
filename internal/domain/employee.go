package domain

import (
	"time"
)

// Employee represents an employee in the system
type Employee struct {
	NIP       int64     `json:"employee_id" db:"nip"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// EmployeeRequest represents a request to create a new employee
type EmployeeRequest struct {
	EmployeeID int64  `json:"employee_id"`
	Name       string `json:"name"`
}

// EmployeeResponse represents a response after creating a new employee
type EmployeeResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

// DeleteEmployeeRequest represents a request to delete an employee
type DeleteEmployeeRequest struct {
	EmployeeID int64 `json:"employee_id"`
}

// DeleteEmployeeResponse represents a response after deleting an employee
type DeleteEmployeeResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}
