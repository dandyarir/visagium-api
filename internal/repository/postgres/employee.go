package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"visagium-api/internal/domain"
	"visagium-api/internal/repository"
)

// EmployeeRepository is a PostgreSQL implementation of repository.EmployeeRepository
type EmployeeRepository struct {
	db *sql.DB
}

// NewEmployeeRepository creates a new PostgreSQL employee repository
func NewEmployeeRepository(db *sql.DB) repository.EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

// Create creates a new employee in the database
func (r *EmployeeRepository) Create(ctx context.Context, employee *domain.Employee) (int64, error) {
	query := `INSERT INTO "Employee" (nip, name, created_at) VALUES ($1, $2, $3) RETURNING nip`

	var nip int64
	err := r.db.QueryRowContext(
		ctx,
		query,
		employee.NIP,
		employee.Name,
		time.Now(),
	).Scan(&nip)

	if err != nil {
		return 0, fmt.Errorf("failed to create employee: %w", err)
	}

	return nip, nil
}

// Delete deletes an employee from the database
func (r *EmployeeRepository) Delete(ctx context.Context, employeeID int64) error {
	query := `UPDATE "Employee" SET deleted_at = NOW() WHERE nip = $1`

	result, err := r.db.ExecContext(ctx, query, employeeID)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("employee with ID %d not found", employeeID)
	}

	return nil
}

// GetByID gets an employee by ID
func (r *EmployeeRepository) GetByID(ctx context.Context, employeeID int64) (*domain.Employee, error) {
	query := `SELECT nip, name, created_at FROM "Employee" WHERE nip = $1 AND deleted_at IS NULL`

	var employee domain.Employee
	err := r.db.QueryRowContext(ctx, query, employeeID).Scan(
		&employee.NIP,
		&employee.Name,
		&employee.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get employee: %w", err)
	}

	return &employee, nil
}
