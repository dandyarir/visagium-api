package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"visagium-api/internal/domain"
	"visagium-api/internal/repository"
)

// AttendanceRepository is a PostgreSQL implementation of repository.AttendanceRepository
type AttendanceRepository struct {
	db *sql.DB
}

// NewAttendanceRepository creates a new PostgreSQL attendance repository
func NewAttendanceRepository(db *sql.DB) repository.AttendanceRepository {
	return &AttendanceRepository{
		db: db,
	}
}

// Create creates a new attendance log in the database
func (r *AttendanceRepository) Create(ctx context.Context, attendance *domain.AttendanceLog) error {
	query := `INSERT INTO "AttendanceLog" (id, employee_nip, activity_type, notes, timestamp)
			  VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		attendance.ID,
		attendance.EmployeeID,
		attendance.ActivityType,
		attendance.Notes,
		attendance.Timestamp,
	)

	if err != nil {
		return fmt.Errorf("failed to create attendance log: %w", err)
	}

	return nil
}

// GetByDate gets attendance logs by date
func (r *AttendanceRepository) GetByDate(ctx context.Context, date string) ([]domain.AttendanceData, error) {
	query := `
		WITH ClockInOut AS (
			SELECT
				e.nip as employee_id,
				e.name,
				MIN(CASE WHEN al.activity_type = 'check-in' THEN TO_CHAR(al.timestamp, 'HH24:MI:SS') END) as clock_in_time,
				MAX(CASE WHEN al.activity_type = 'check-out' THEN TO_CHAR(al.timestamp, 'HH24:MI:SS') END) as clock_out_time
			FROM
				"Employee" e
			JOIN
				"AttendanceLog" al ON e.nip = al.employee_nip
			WHERE
				DATE(al.timestamp) = $1
			GROUP BY
				e.nip, e.name
		)
		SELECT
			employee_id,
			name,
			clock_in_time,
			clock_out_time
		FROM
			ClockInOut
		ORDER BY
			employee_id
	`

	rows, err := r.db.QueryContext(ctx, query, date)
	if err != nil {
		return nil, fmt.Errorf("failed to query attendance logs: %w", err)
	}
	defer rows.Close()

	var attendances []domain.AttendanceData
	for rows.Next() {
		var a domain.AttendanceData
		var clockOutTime sql.NullString

		if err := rows.Scan(&a.EmployeeID, &a.Name, &a.ClockInTime, &clockOutTime); err != nil {
			return nil, fmt.Errorf("failed to scan attendance row: %w", err)
		}

		if clockOutTime.Valid {
			clockOutStr := clockOutTime.String
			a.ClockOutTime = &clockOutStr
		} else {
			a.ClockOutTime = nil
		}

		attendances = append(attendances, a)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating attendance rows: %w", err)
	}

	return attendances, nil
}
