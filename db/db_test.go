package db_test

import (
	"testing"

	"github.com/chandan782/Pismo/db"
	"github.com/stretchr/testify/assert"
)

func TestInitDBSuccess(t *testing.T) {
	setupDB(t)
	assert.NotNil(t, db.DB)
}

func TestCloseDB(t *testing.T) {
	setupDB(t)
	db.CloseDB()

	var result interface{}
	err := db.DB.Raw("SELECT 1").Scan(&result).Error
	if err != nil && err.Error() == "sql: database is closed" {
		t.Log("Database connection closed successfully")
	} else {
		t.Errorf("Expected 'sql: database is closed' error, got: %v", err)
	}
}
