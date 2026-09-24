E-Commerce Backend

A backend service for an e-commerce application built with Go and PostgreSQL, following a layered and domain-oriented project structure.

The project is designed to demonstrate how a production-style Go backend can be organized with separation of concerns between domain logic, repositories, infrastructure, HTTP/REST handlers, database migrations, and configuration.

🚀 Tech Stack
Go 1.25+
PostgreSQL
SQLX — Database access
SQL Migrate — Database migrations
JWT — Authentication
UUID — Unique entity identifiers
Godotenv — Environment configuration
✨ Features
RESTful backend architecture
User management
Product management
PostgreSQL database integration
Database migrations
JWT-based authentication configuration
Environment-based configuration
Repository pattern for data access
Domain-oriented application structure
Separation of business logic and infrastructure
UUID-based entity identification
🏗️ Project Structure
ecommerce/
│
├── cmd/                 # Application/server startup
│
├── config/              # Application configuration
│
├── db_Query/            # Database queries
│
├── domain/              # Core domain models and business logic
│
├── infra/
│   └── db/              # Database infrastructure
│
├── migrations/          # Database migration files
│
├── product/             # Product domain and functionality
│
├── repo/                # Repository/data-access layer
│
├── rest/                # REST API layer
│
├── user/                # User domain and functionality
│
├── utils/               # Shared utility functions
│
├── main.go              # Application entry point
├── go.mod               # Go module definition
└── go.sum               # Dependency checksums
🧠 Architecture

The application follows a layered architecture where different responsibilities are separated into dedicated packages.

                ┌─────────────────────┐
                │      REST API       │
                │       /rest         │
                └──────────┬──────────┘
                           │
                           ▼
                ┌─────────────────────┐
                │     Domain Logic    │
                │  product / user     │
                └──────────┬──────────┘
                           │
                           ▼
                ┌─────────────────────┐
                │     Repository      │
                │       /repo         │
                └──────────┬──────────┘
                           │
                           ▼
                ┌─────────────────────┐
                │     PostgreSQL      │
                │      /infra/db      │
                └─────────────────────┘

This separation makes the application easier to maintain, test, and extend.

⚙️ Requirements

Before running the project, make sure you have:

Go 1.25 or later
PostgreSQL
Git
📥 Installation

Clone the repository:

git clone https://github.com/developerPranRoy/ecommerce.git

Navigate into the project:

cd ecommerce

Download dependencies:

go mod download
🔐 Environment Variables

The application loads configuration from a .env file.

Create a .env file in the project root:

VERSION=1.0.0
SERVICE_NAME=ecommerce
HTTP_PORT=8080

JWT_SECRET_KEY=your-secret-key

DB_HOST=localhost
DB_PORT=5432
DB_NAME=ecommerce
DB_USER=postgres
DB_PASSWORD=your-password
DB_ENABLE_SSL_MODE=false
Configuration
Variable	Description
VERSION	Application version
SERVICE_NAME	Service name
HTTP_PORT	HTTP server port
JWT_SECRET_KEY	Secret key used for JWT
DB_HOST	PostgreSQL host
DB_PORT	PostgreSQL port
DB_NAME	Database name
DB_USER	Database username
DB_PASSWORD	Database password
DB_ENABLE_SSL_MODE	PostgreSQL SSL mode

The application validates the required environment variables during startup.

🗄️ Database Setup

Create a PostgreSQL database:

CREATE DATABASE ecommerce;

Configure the database credentials in .env.

Run the project's database migrations according to the migration files inside:

migrations/

The project uses sql-migrate for database migrations and sqlx for database interaction.

▶️ Running the Application

Run the application with:

go run main.go

The application entry point calls the server functionality from the cmd package.

For development, you can also build the application:

go build -o ecommerce

Then run:

./ecommerce
🧪 Testing

Run all Go tests:

go test ./...

Run tests with additional output:

go test -v ./...
📡 API

The application exposes functionality through a REST API.

The main REST layer is located inside:

rest/

Domain-specific functionality is separated into areas such as:

user/
product/

This structure makes it easier to add additional modules such as:

order/
cart/
payment/
category/
review/

in the future.

🔒 Security

The project includes JWT configuration for authentication.

For production environments:

Never commit .env files.
Use a strong JWT secret.
Use HTTPS.
Use a strong PostgreSQL password.
Enable PostgreSQL SSL when required.
Validate and sanitize incoming API data.
Keep secrets outside the source code.
🛠️ Development

Install dependencies:

go mod download

Format the code:

go fmt ./...

Run static checks:

go vet ./...

Run tests:

go test ./...
📈 Future Improvements

Potential improvements for the project include:

Complete authentication and authorization flow

Product CRUD APIs

Cart management

Order management

Payment integration

Product search and filtering

Pagination

Request validation

API documentation with Swagger/OpenAPI

Unit and integration tests

Docker support

Docker Compose for PostgreSQL

Structured logging

Redis caching

Rate limiting

CI/CD pipeline

Observability and monitoring

📚 Project Goals

This project focuses on learning and applying backend engineering concepts using Go, including:

Clean separation of concerns
Domain-driven organization
Repository patterns
REST API development
PostgreSQL database design
Database migrations
Authentication
Configuration management
Maintainable backend architecture
👨‍💻 Author

Pran Kumar Roy

Full-Stack Software Developer

GitHub: developerPranRoy
