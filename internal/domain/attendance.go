package domain

import (
	"time"
)

// AttendanceLog represents an attendance log entry
type AttendanceLog struct {
	ID           string    `json:"id" db:"id"`
	EmployeeID   int64     `json:"employee_id" db:"employee_nip"`
	ActivityType string    `json:"activity_type" db:"activity_type"`
	Notes        string    `json:"notes" db:"notes"`
	Timestamp    time.Time `json:"timestamp" db:"timestamp"`
}

// AttendanceRequest represents a request to submit attendance
type AttendanceRequest struct {
	EmployeeID int64  `json:"employee_id"`
	Name       string `json:"name"`
	Timestamp  string `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

// AttendanceResponse represents a response after submitting attendance
type AttendanceResponse struct {
	Message string `json:"message"`
}

// AttendanceData represents the data structure for a single attendance record
type AttendanceData struct {
	EmployeeID   int64   `json:"employe_id"`
	Name         string  `json:"name"`
	ClockInTime  *string  `json:"clock_in_time"`
	ClockOutTime *string `json:"clock_out_time"`
}

// AttendanceListResponse represents the response for attendance list
type AttendanceListResponse struct {
	Date        string           `json:"date"`
	Count       int              `json:"count"`
	Attendances []AttendanceData `json:"attendaces"`
}
