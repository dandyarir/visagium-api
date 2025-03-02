# Visagium Attendance API

## Overview
This is a RESTful API for employee attendance management built with Go and Echo framework. The API allows managing employees and recording attendance logs.

## Features
- Employee management (create, delete)
- Attendance logging (clock in/out)
- Attendance reports by date

## Technology Stack
- Go (Golang)
- Echo Framework
- PostgreSQL
- Docker

## Getting Started

### Prerequisites
- Go 1.21+
- Docker and Docker Compose (for containerized setup)
- PostgreSQL (if running locally)

### Installation and Setup

1. Clone the repository
```bash
git clone https://github.com/yourusername/visagium-api.git
cd visagium-api
```

2. Set up environment variables (copy from `.env.sample`)
```bash
cp .env.sample .env
# Edit .env with your configuration
```

3. Run with Docker Compose
```bash
docker-compose up -d
```

Or build and run locally:
```bash
go mod download
go run cmd/api/main.go
```

## API Endpoints

### Employee Management
- `POST /Employee` - Register a new employee
- `DELETE /Employee` - Delete an employee

### Attendance Management
- `POST /Attendance` - Submit attendance log
- `GET /Attendance?date=YYYY-MM-DD` - Retrieve attendance logs for a specific date

## Project Structure
```
visagium-api/
├── cmd/                    # Application entry points
│   └── api/                # API server
├── internal/               # Private application code
│   ├── config/             # Configuration
│   ├── domain/             # Domain models
│   ├── repository/         # Data access layer
│   ├── service/            # Business logic
│   ├── handler/            # HTTP handlers
│   └── router/             # Routes setup
├── pkg/                    # Public libraries
│   └── database/           # Database utilities
```

## License
MIT
