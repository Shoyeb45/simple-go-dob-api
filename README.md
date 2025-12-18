# Simple Date of Birth App
![Go](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Zap](https://img.shields.io/badge/Zap-Logger-orange?style=for-the-badge)

A clean and efficient backend application written in Go for managing user data with date of birth tracking.

### The backend service is deployed
### Live URL: https://simple-go-dob-api.onrender.com/

## Features

The application provides a complete REST API with the following endpoints:

- **Create User** - `POST /users` - Register a new user
- **Get All Users** - `GET /users` - Retrieve list of all users
- **Get User** - `GET /users/:id` - Fetch details of a specific user
- **Update User** - `PUT /users/:id` - Update existing user information
- **Delete User** - `DELETE /users/:id` - Remove a user from the system

## Tech Stack


- **Language:** Go 1.25.5
- **Web Framework:** [Fiber](https://gofiber.io/) - Express-inspired web framework
- **Database:** PostgreSQL with SQLC for type-safe queries
- **Logging:** [Uber Zap](https://github.com/uber-go/zap) - High-performance structured logging
- **Validation:** [go-playground/validator](https://github.com/go-playground/validator) - Struct validation
- **Migration:** [golang-migrate](https://github.com/golang-migrate/migrate) - Database schema management

## Project Structure

```
.
├── cmd
│   └── server          # Application entry point
├── config              # Configuration and environment variables
├── db                  # Database layer
│   ├── migrations      # SQL migration files
│   └── sqlc            # Generated type-safe Go code from SQL
└── internal
    ├── app             # Application setup and initialization
    ├── core            # Utility functions and helpers
    ├── database        # PostgreSQL connection management
    ├── handler         # HTTP request handlers
    ├── logger          # Logging configuration
    ├── mapper          # JSON response transformers
    ├── middlewares     # HTTP middlewares (auth, logging, etc.)
    ├── models          # DTOs and request/response schemas
    ├── repository      # Data access layer
    ├── routes          # Route definitions
    └── service         # Business logic layer
```

## Architecture

<p align="center">
    <img src="./public/architecture.png" alt="Architecture Diagram">
    <br>
    <em>High-level application architecture</em>
</p>

The application follows a clean architecture pattern with clear separation of concerns:

- **Handler Layer** - Handles HTTP requests and responses
- **Service Layer** - Contains business logic
- **Repository Layer** - Manages data access and persistence
- **Database Layer** - PostgreSQL with type-safe SQLC queries

## Getting Started

### Prerequisites

- Go version >= 1.25.5
- PostgreSQL database
- Make (optional, for using Makefile commands)

### Installation

1. **Install SQLC** (for generating type-safe database code):
   ```bash
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   ```

2. **Install golang-migrate** (for database migrations):
   
   **macOS:**
   ```bash
   brew install golang-migrate
   ```
   
   **Linux:**
   ```bash
   curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
   sudo mv migrate /usr/local/bin
   ```
   
   **Windows:**
   ```bash
   # Download from: https://github.com/golang-migrate/migrate/releases
   ```

### Setup

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd simple-dob-app
   ```

2. **Generate SQLC code** (skip if already generated):
   ```bash
   sqlc generate
   ```

3. **Create initial migration** (skip if already exists):
   ```bash
   migrate create -ext sql -dir db/migrations -seq init
   ```

4. **Configure environment variables:**
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` and fill in your configuration:
   ```env
   DATABASE_URL=postgresql://user:password@localhost:5432/dbname?sslmode=disable
   PORT=3000
   LOG_LEVEL=info
   ```

5. **Run database migrations:**
   ```bash
   migrate -path ./db/migrations/ -database "postgresql://user:password@localhost:5432/dbname?sslmode=disable" up
   ```

6. **Install dependencies:**
   ```bash
   go mod download
   ```

7. **Run the application:**
   ```bash
   # Development mode
   go run cmd/server/main.go
   
   # Build and run
   go build -o bin/server cmd/server/main.go
   ./bin/server
   ```

The server will start on `http://localhost:3000` (or the port specified in your `.env` file).

## API Documentation

Once the server is running, you can test the endpoints:

```bash
# Create a user
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","date_of_birth":"1990-01-15"}'

# Get all users
curl http://localhost:3000/users

# Get specific user
curl http://localhost:3000/users/1

# Update user
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","date_of_birth":"1990-01-15"}'

# Delete user
curl -X DELETE http://localhost:3000/users/1
```

## Running Tests

```bash
# Run Unit test just for CalculateAge function
go test ./tests
```

## Building for Production

> Build directly 

```bash
# Build binary
go build -o bin/server cmd/server/main.go

# Build with optimizations
go build -ldflags="-s -w" -o bin/server cmd/server/main.go
```

> Build using Docker

```bash
# Build docker image
docker build -t my-go-lang-app .

# Run the image
docker run \
   -p 8080:8080 \ # port mapping
   --env-file \ # provide env variables
   my-go-lang-app
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.



## 👤 Author

Shoyeb Ansari - [@Shoyeb45](https://github.com/Shoyeb45)

## 🙏 Acknowledgments

- [Alex Mux](https://youtu)
- [Fiber Framework](https://gofiber.io/)
- [SQLC](https://sqlc.dev/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Uber Zap](https://github.com/uber-go/zap)