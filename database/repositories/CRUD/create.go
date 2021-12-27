package CRUD

import (
	"Golang_Fiber/database"
	"Golang_Fiber/messages"
	"Golang_Fiber/utilities"
	"errors"
	"gorm.io/gorm"
)

func Create(data interface{}) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Create(data).Error
}

func CreateWithTransaction(data interface{}, tx *gorm.DB) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return tx.Create(data).Error
}

func CreateWithoutAssociations(data interface{}, association string) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return database.DB.Omit(association).Create(data).Error
}

func CreateWithTransactionWithoutAssociations(data interface{}, association string, tx *gorm.DB) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	return tx.Omit(association).Create(data).Error
}
