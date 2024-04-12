# Pismo

## Description
This is a Go application for managing accounts and transactions.

## Prerequisites
- Go installed on your machine
- PostgreSQL database

## Getting Started
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/pismo.git
   cd pismo


### Set up the required environment variables:
export DB_HOST="your_db_host"
export DB_PORT=5432
export DB_USER="your_db_user"
export DB_PASSWORD="your_db_password"
export DB_NAME="your_db_name"
export DB_SSL_MODE="disable"
export SERVER_PORT="3000"

### Generate Swagger documentation:
swag init

### Generate mocks:
mockgen -source=./db/db_handler.go -destination=./tests/mocks/db_mocks.go -package=mocks
mockgen -source=./validate/validate.go -destination=./tests/mocks/validate_mocks.go -package=mocks
mockgen -source=./api/services/services.go -destination=./tests/mocks/services_mocks.go -package=mocks
