package db

import (
	"gorm.io/gorm"
)

type DBHandlerInterface interface {
	Create(model interface{}) error
	ReadByID(model interface{}, query interface{}, args ...interface{}) error
	ReadAll(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}, id interface{}) error
}

type DBHandler struct {
	DB *gorm.DB
}

// create a new DBHandler instance
func NewDBHandler(db *gorm.DB) DBHandlerInterface {
	return &DBHandler{DB: db}
}

// create operation
func (h *DBHandler) Create(model interface{}) error {
	return h.DB.Create(model).Error
}

// read operation - get by ID
func (h *DBHandler) ReadByID(model interface{}, query interface{}, args ...interface{}) error {
	return h.DB.First(model, query, args).Error
}

// read operation - get all records
func (h *DBHandler) ReadAll(model interface{}) error {
	return h.DB.Find(model).Error
}

// update operation
func (h *DBHandler) Update(model interface{}) error {
	return h.DB.Save(model).Error
}

// delete operation
func (h *DBHandler) Delete(model interface{}, id interface{}) error {
	return h.DB.Delete(model, id).Error
}
