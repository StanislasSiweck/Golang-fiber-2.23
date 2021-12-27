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
