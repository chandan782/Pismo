package db_test

import (
	"testing"

	"github.com/chandan782/Pismo/db"
	dot "github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type User struct {
	ID   uint   `gorm:"id"`
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`
}

func setupDB(t *testing.T) {
	err := dot.Load(".env.sample")
	if err != nil {
		t.Fatalf("failed to load envs: %v", err)
	}

	// db initialization
	db.InitDB()

	// auto-migrate the models
	err = db.DB.AutoMigrate(&User{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}
}

func TestCreate(t *testing.T) {
	// setup DB
	setupDB(t)
	handler := db.NewDBHandler(db.DB)
	defer db.CloseDB()

	// create operation test
	user := User{Name: "John", Age: 30}
	err := handler.Create(&user)
	assert.NoError(t, err, "create operation should not return error")
	assert.NotEqual(t, 0, user.ID, "user ID should be set after creation")
}

func TestReadByID(t *testing.T) {
	// setup DB
	setupDB(t)
	handler := db.NewDBHandler(db.DB)
	defer db.CloseDB()

	// create a user to read
	user := User{Name: "John", Age: 30}
	err := handler.Create(&user)
	assert.NoError(t, err, "create operation should not return error")

	// read by ID operation test
	var readUser User
	err = handler.ReadByID(&readUser, "id = ?", user.ID)
	assert.NoError(t, err, "ReadByID operation should not return error")
	assert.Equal(t, user.Name, readUser.Name, "user names should match")
}

func TestReadAll(t *testing.T) {
	// setup DB
	setupDB(t)
	handler := db.NewDBHandler(db.DB)
	defer db.CloseDB()

	// read all operation test
	var users []User
	err := handler.ReadAll(&users)
	assert.NoError(t, err, "read all operation should not return error")
}

func TestUpdate(t *testing.T) {
	// setup DB
	setupDB(t)
	handler := db.NewDBHandler(db.DB)
	defer db.CloseDB()

	// create a user to update
	user := User{Name: "John", Age: 30}
	err := handler.Create(&user)
	assert.NoError(t, err, "create operation should not return error")

	// update operation test
	user.Name = "Updated Name"
	err = handler.Update(&user)
	assert.NoError(t, err, "update operation should not return error")
}

func TestDelete(t *testing.T) {
	// setup DB
	setupDB(t)
	handler := db.NewDBHandler(db.DB)
	defer db.CloseDB()

	// create a user to delete
	user := User{Name: "John", Age: 30}
	err := handler.Create(&user)
	assert.NoError(t, err, "create operation should not return error")

	// delete operation test
	err = handler.Delete(&user, user.ID)
	assert.NoError(t, err, "delete operation should not return error")
}
