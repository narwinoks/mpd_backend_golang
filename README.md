# SIMRS Enterprise Backend

Modular Clean Architecture Backend for Hospital Information Management System (SIMRS) built with Golang.

## 🚀 Project Information
- **Package Name:** `backend-app`
- **Version:** `1.0.0`
- **Framework:** Gin Gonic
- **ORM:** GORM (PostgreSQL)
- **Architecture:** Modular Domain Driven Design (DDD)

## 📋 Prerequisites
Ensure you have the following installed on your system:
- **Go:** 1.21 or later
- **PostgreSQL:** 14 or later
- **Wire CLI:** For Dependency Injection
  ```bash
  go install github.com/google/wire/cmd/wire@latest
  ```
- **Migrate CLI:** For Database Migrations
  ```bash
  # Windows (via Scoop)
  scoop install golang-migrate
  # macOS
  brew install golang-migrate
  ```
- **Air CLI:** For Hot Reloading
  ```bash
  go install github.com/air-verse/air@latest
  ```

## 🛠️ Setup Instructions

### 1. Clone & Install Dependencies
```bash
git clone <repository-url>
cd backend-app
go mod tidy
```

### 2. Configuration
Copy the configuration template and update your database credentials:
- Open `config/config.yaml`
- Update the `database` section:
```yaml
database:
  host: "localhost"
  port: 5432
  user: "your_user"
  password: "your_password"
  name: "simrs_db"
  sslmode: "disable"
```

### 3. Database Migration & Seeding
Run the following command to create the initial tables:
```bash
migrate -path db/migrations -database "postgres://postgres:password@localhost:5432/simrs_db?sslmode=disable" up
```

To seed the database with initial/dummy data:
```bash
go run cmd/seeder/main.go
```

### 4. Dependency Injection
Generate the wire injector code before running the app:
```bash
wire ./internal/modules/master
```

## 🏃 Running the Project

### Development Mode (with Hot Reload)
Recommended for development. The app will automatically rebuild and restart on file changes.
```bash
air
```

### Production Mode
```bash
go run cmd/api/main.go
```

## 📂 Project Structure
```text
backend-app/
├── cmd/api/                # Entry point
├── config/                 # Viper & Logrus setup
├── db/migrations/          # SQL migration files
├── internal/
│   ├── core/               # Shared utilities (DB, Response, Middleware)
│   └── modules/            # Domain modules
│       └── master/         # Master Data Module
│           ├── controller/ # Gin Handlers
│           ├── model/      # GORM Entities
│           ├── repository/ # Data access layer
│           ├── service/    # Business logic layer
│           ├── request/    # Validation DTOs
│           └── response/   # API response DTOs
```

## 📡 API Endpoints (Master Module)
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/v1/master/users/` | Fetch all users |
| `GET` | `/api/v1/master/users/:id` | Fetch user by ID |
| `POST` | `/api/v1/master/users/` | Create a new user |
