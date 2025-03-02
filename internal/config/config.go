package config

import (
	"fmt"
	"os"
)

// Config holds application configuration
type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Server
	ServerAddress string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "postgres"),
		DBName:        getEnv("DB_NAME", "visagium_backend"),
		DBSSLMode:     getEnv("DB_SSL_MODE", "disable"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
	}

	return cfg, nil
}

// PostgresConnectionString returns the connection string for PostgreSQL
func (c *Config) PostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
