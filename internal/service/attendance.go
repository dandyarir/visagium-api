package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"visagium-api/internal/domain"
	"visagium-api/internal/repository"
)

// AttendanceServiceImpl is an implementation of service.AttendanceServiceImpl
type AttendanceServiceImpl struct {
	attendanceRepo repository.AttendanceRepository
	employeeRepo   repository.EmployeeRepository
}

// NewAttendanceService creates a new AttendanceService
func NewAttendanceService(attendanceRepo repository.AttendanceRepository, employeeRepo repository.EmployeeRepository) *AttendanceServiceImpl {
	return &AttendanceServiceImpl{
		attendanceRepo: attendanceRepo,
		employeeRepo:   employeeRepo,
	}
}

// SubmitAttendance submits an attendance log
func (s *AttendanceServiceImpl) SubmitAttendance(ctx context.Context, req *domain.AttendanceRequest) (*domain.AttendanceResponse, error) {
	// Verify employee exists
	employee, err := s.employeeRepo.GetByID(ctx, req.EmployeeID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify employee: %w", err)
	}

	if employee == nil {
		return nil, fmt.Errorf("employee with ID %d not found", req.EmployeeID)
	}

	// Determine activity type based on time of day (simplified)
	timeStamp, err := time.Parse("2006-01-02 15:04:05", req.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse timestamp: %w", err)
	}
	activityType := "check-in"
	hour := timeStamp.Hour()
	if hour >= 12 { // After noon, consider it clock out
		activityType = "check-out"
	}

	// Create attendance log
	attendance := &domain.AttendanceLog{
		ID:           uuid.New().String(),
		EmployeeID:   req.EmployeeID,
		ActivityType: activityType,
		Notes:        "",
		Timestamp:    timeStamp,
	}

	err = s.attendanceRepo.Create(ctx, attendance)
	if err != nil {
		return nil, fmt.Errorf("failed to submit attendance: %w", err)
	}

	return &domain.AttendanceResponse{
		Message: "data absensi tersimpan",
	}, nil
}

// GetAttendanceByDate gets attendance logs by date
func (s *AttendanceServiceImpl) GetAttendanceByDate(ctx context.Context, date string) (*domain.AttendanceListResponse, error) {
	attendances, err := s.attendanceRepo.GetByDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance data: %w", err)
	}

	groupedAttendaces := make(map[int64]domain.AttendanceData)
	for _, attendance := range attendances {
		var (
			clockIn  string
			clockOut *string
		)
		var atd domain.AttendanceData
		if _, exists := groupedAttendaces[attendance.EmployeeID]; !exists {
			atd = domain.AttendanceData{
				EmployeeID: attendance.EmployeeID,
				Name:       attendance.Name,
			}
		} else {
			atd = groupedAttendaces[attendance.EmployeeID]
		}

		if attendance.ClockInTime != nil {
			clockIn = *attendance.ClockInTime
		}
		if attendance.ClockOutTime != nil {
			clockOut = attendance.ClockOutTime
		}

		atd.ClockInTime = &clockIn
		atd.ClockOutTime = clockOut

		groupedAttendaces[attendance.EmployeeID] = atd
	}

	attendancesRes := make([]domain.AttendanceData, 0, len(groupedAttendaces))
	for _, atd := range groupedAttendaces {
		attendancesRes = append(attendancesRes, atd)
	}

	return &domain.AttendanceListResponse{
		Date:        date,
		Count:       len(attendancesRes),
		Attendances: attendancesRes,
	}, nil
}
