package CRUD

import (
	"Golang_Fiber/database"
	"Golang_Fiber/messages"
	"Golang_Fiber/utilities"
	"errors"
)

func GetOne(data interface{}, id uint) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Model(data).Limit(1).Find(data, id).Error
}

func GetOneByColumn(data interface{}, column string, value interface{}, joins ...string) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	statement := database.DB.Model(data)
	return statement.Limit(1).Find(data, column, value).Error
}
