package CRUD

import (
	"Golang_Fiber/database"
	"Golang_Fiber/messages"
	"Golang_Fiber/utilities"
	"errors"
	"gorm.io/gorm"
)

func Delete(object interface{}) error {
	if utilities.IsStruct(object) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Delete(object).Error
}

func DeleteWithTransaction(object interface{}, tx *gorm.DB) error {
	if utilities.IsStruct(object) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return tx.Delete(object).Error
}

func ForceDelete(object interface{}) error {
	if utilities.IsStruct(object) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Unscoped().Delete(object).Error
}

func ForceDeleteWithTransaction(object interface{}, tx *gorm.DB) error {
	if utilities.IsStruct(object) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return tx.Unscoped().Delete(object).Error
}
