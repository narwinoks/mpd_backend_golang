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
wire ./internal/modules/auth
wire ./internal/modules/master
```

## ⚡ Caching (Redis & In-Memory)
This project features a unified cache store that supports both Redis and In-Memory fallback.

### Redis Configuration
Update `config/config.yaml` to enable Redis:
```yaml
redis:
  enabled: true # Set to true to use Redis
  host: "localhost"
  port: 6379
  password: ""
  db: 0
```
If `enabled` is `false` or the connection fails, the system automatically falls back to **In-Memory** caching.

## 🛡️ Authorization Middlewares
We provide two specialized middlewares for granular access control, both utilizing the caching system for optimal performance.

### 1. Permission Middleware
Used to check for specific action permissions (stored in `app_permission_m`).

**Example Usage in Router:**
```go
user.GET("/sensitive-data", 
    middleware.AuthMiddleware(r.config, r.tokenRepo), 
    middleware.PermissionMiddleware(r.permRepo, "USER_VIEW"), 
    r.userController.GetProfile)
```

### 2. Module Middleware
Used to check if a user has access to a specific application module (stored in `app_modules_m`).

**Example Usage in Router:**
```go
user.GET("/admin-panel", 
    middleware.AuthMiddleware(r.config, r.tokenRepo), 
    middleware.ModuleMiddleware(r.moduleRepo, "MOD_ADMIN_MANAGEMENT"), 
    r.userController.GetProfile)
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

## 🧪 Running Unit Tests
This project uses `github.com/stretchr/testify` and `github.com/brianvoe/gofakeit/v6` for unit testing. Tests follow the rules defined in `ai/test.md`.

### Run all tests
```bash
go test ./...
```

### Run specific module tests
```bash
go test ./internal/modules/master/test/...
```

### Run with coverage
```bash
go test -cover ./...
```

## 📝 Commit Message Rules
We follow the **[Conventional Commits](https://www.conventionalcommits.org/)** specification. `husky` and `commitlint` are used to enforce these rules.

### ⚙️ Setup Hooks
Ensure Node.js dependencies are installed to activate the git hooks:
```bash
npm install
```

### 📋 Format
Commit messages MUST follow this format:
`<type>(<optional scope>): <description>`

**Common types:**
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

**Example:**
`feat(auth): add user login functionality`

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
