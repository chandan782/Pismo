package configs

import (
	"os"
	"strconv"
)

// DBConfig holds database configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// GetDBConfig returns DBConfig loaded from environment variables
func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}

// ServerConfig holds server related configurations
type ServerConfig struct {
	Port string
}

// GetServerConfig returns ServerConfig loaded from environment variables
func GetServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}
}

// getEnvAsInt converts environment variable to int
func getEnvAsInt(key string, defaultValue int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultValue
	}
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}
	return valInt
}
