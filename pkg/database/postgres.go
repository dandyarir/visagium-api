package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"visagium-api/internal/config"
)

// NewPostgresConnection creates a new connection to PostgreSQL database
func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.PostgresConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
