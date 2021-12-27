package CRUD

import (
	"Golang_Fiber/database"
	"Golang_Fiber/messages"
	"Golang_Fiber/utilities"
	"errors"
	"gorm.io/gorm"
)

func Update(data interface{}) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Model(data).Updates(data).Error
}

func UpdateWithTransaction(data interface{}, tx *gorm.DB) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return tx.Model(data).Updates(data).Error
}

func UpdateSpecificField(data interface{}, field string, value interface{}) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Model(data).Update(field, value).Error
}

func UpdateSpecificFieldWithTransaction(data interface{}, field string, value interface{}, tx *gorm.DB) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return tx.Model(data).Update(field, value).Error
}

func RestoreSoftDelete(data interface{}) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Model(data).Set("is_restore", true).Update("deleted_at", nil).Error
}
func RestoreSoftDeleteWithTransaction(data interface{}, tx *gorm.DB) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return tx.Model(data).Update("deleted_at", nil).Error
}
