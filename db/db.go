package db

import (
	"fmt"
	"log"

	"github.com/chandan782/Pismo/configs"
	"github.com/chandan782/Pismo/db/schemas"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	conf := configs.GetDBConfig()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName, conf.SSLMode)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.Errorf("error connecting to database: %v", err)
	}

	migrateDBSchemas(DB)

	return nil
}

func migrateDBSchemas(db *gorm.DB) {
	db.AutoMigrate(&schemas.Account{})
	db.AutoMigrate(&schemas.Transaction{})
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("error getting DB instance: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("error closing database: %v", err)
	}
}
