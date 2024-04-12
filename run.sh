#!/bin/bash

# db required env variables
export DB_HOST="kiouni.db.elephantsql.com"
export DB_PORT=5432
export DB_USER="rghrxzte"
export DB_PASSWORD="5uzggMvXHjpT2XaUpsOQzKfS37BBqMlz"
export DB_NAME="rghrxzte"
export DB_SSL_MODE="disable"

# server required env variable
export SERVER_PORT="3000"

# generate swagger docs
swag init

# mockgen command to generates mocks
mockgen -source=./db/db_handler.go -destination=./tests/mocks/db_mocks.go -package=mocks
mockgen -source=./validate/validate.go -destination=./tests/mocks/validate_mocks.go -package=mocks
mockgen -source=./api/services/services.go -destination=./tests/mocks/services_mocks.go -package=mocks

# refresh dependencies
go mod tidy

# run the Go application
go run main.go
