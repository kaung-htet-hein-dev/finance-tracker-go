# Finance Tracker Go

A RESTful API for personal finance tracking built with Go (Echo framework), featuring JWT authentication, CRUD operations for transactions, and AI-powered financial insights.

## Features

- **User Authentication**: JWT-based registration and login system
- **Transaction Management**: Full CRUD operations for financial transactions
- **AI Insights**: Intelligent financial analytics and recommendations
- **Clean Architecture**: Organized with handlers, services, and models
- **Input Validation**: Comprehensive request validation and error handling
- **Database Support**: SQLite (with easy PostgreSQL migration path)
- **Containerized**: Docker support for easy deployment

## Tech Stack

- **Framework**: Echo v4
- **ORM**: GORM
- **Database**: SQLite (PostgreSQL ready)
- **Authentication**: JWT with bcrypt password hashing
- **Validation**: go-playground/validator
- **Containerization**: Docker & Docker Compose

## Project Structure

```
├── cmd/server/          # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── database/        # Database connection and migrations
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/      # Custom middleware (JWT auth)
│   ├── models/          # Data models and DTOs
│   └── services/        # Business logic layer
├── Dockerfile           # Container configuration
├── docker-compose.yml   # Local development setup
└── README.md
```

## Quick Start

### Using Docker (Recommended)

1. Clone the repository:
```bash
git clone <repository-url>
cd finance-tracker-go
```

2. Start with Docker Compose:
```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`

### Manual Setup

1. Install Go 1.21 or later
2. Clone and navigate to the project:
```bash
git clone <repository-url>
cd finance-tracker-go
```

3. Install dependencies:
```bash
go mod download
```

4. Run the application:
```bash
go run cmd/server/main.go
```

## API Endpoints

### Authentication

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/auth/register` | Register a new user |
| POST | `/api/v1/auth/login` | Login user |

### Transactions (Protected)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/transactions` | Create a new transaction |
| GET | `/api/v1/transactions` | Get user's transactions (paginated) |
| GET | `/api/v1/transactions/:id` | Get specific transaction |
| PUT | `/api/v1/transactions/:id` | Update transaction |
| DELETE | `/api/v1/transactions/:id` | Delete transaction |

### Insights (Protected)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/insights` | Get AI-powered financial insights |

### Health Check

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |

## API Usage Examples

### Register User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "password": "securepassword123"
  }'
```

### Create Transaction
```bash
curl -X POST http://localhost:8080/api/v1/transactions \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "expense",
    "amount": 50.75,
    "category": "groceries",
    "description": "Weekly grocery shopping",
    "date": "2024-01-15T10:30:00Z"
  }'
```

### Get Insights
```bash
curl -X GET http://localhost:8080/api/v1/insights \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Configuration

Environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `DB_TYPE` | `sqlite` | Database type (sqlite/postgres) |
| `DB_NAME` | `finance_tracker.db` | Database name/file |
| `DB_HOST` | `localhost` | Database host (PostgreSQL) |
| `DB_PORT` | `5432` | Database port (PostgreSQL) |
| `DB_USER` | `` | Database user (PostgreSQL) |
| `DB_PASSWORD` | `` | Database password (PostgreSQL) |
| `JWT_SECRET` | `your-secret-key` | JWT signing secret |
| `JWT_EXPIRATION` | `24` | JWT expiration in hours |

## AI Insights Features

The insights endpoint provides:

- **Financial Overview**: Total income, expenses, and net income
- **Category Analysis**: Top spending categories with amounts and transaction counts
- **Monthly Trends**: 6-month income and expense trends
- **Savings Rate**: Percentage of income saved
- **Smart Recommendations**: AI-generated financial advice based on:
  - Spending patterns
  - Savings rate analysis
  - Category-wise expense distribution
  - Income vs expense ratios
  - Emergency fund recommendations

## Development

### Building
```bash
go build -o bin/server cmd/server/main.go
```

### Running Tests
```bash
go test ./...
```

### Database Migrations

Migrations run automatically on startup. The application will create the necessary tables if they don't exist.

## Security Features

- **Password Hashing**: bcrypt with salt
- **JWT Authentication**: Secure token-based auth
- **Input Validation**: Comprehensive request validation
- **SQL Injection Protection**: GORM ORM prevents SQL injection
- **CORS Support**: Configurable cross-origin requests

## License

MIT License - see LICENSE file for details.