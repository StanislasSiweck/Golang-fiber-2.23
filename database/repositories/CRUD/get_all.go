package CRUD

import (
	"Golang_Fiber/database"
	"Golang_Fiber/messages"
	"Golang_Fiber/utilities"
	"errors"
)

func GetAll(data interface{}) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Find(data).Error
}

func GetAllByColumn(data interface{}, column string, value interface{}) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Find(data, column, value).Error
}

func GetAllFindByColumn(data interface{}, column string, value string) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	statement := database.DB.Where(column + " LIKE '%" + value + "%'")
	return statement.Find(data).Error
}
