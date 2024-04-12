#!/bin/bash

# db required env variables
export DB_HOST="<host_name>"
export DB_PORT=5432
export DB_USER="<db_user>"
export DB_PASSWORD="<db_password>"
export DB_NAME="<db_name>"
export DB_SSL_MODE="disable"

# server required env variable
export SERVER_PORT="3000"

# generate swagger docs
swag init -o swagger/docs

# mockgen command to generates mocks
mockgen -source=./db/db_handler.go -destination=./internals/mocks/db_mocks.go -package=mocks
mockgen -source=./validate/validate.go -destination=./internals/mocks/validate_mocks.go -package=mocks
mockgen -source=./api/services/services.go -destination=./internals/mocks/services_mocks.go -package=mocks

# refresh dependencies
go mod tidy

# run the Go application
go run main.go
