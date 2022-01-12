package CRUD

import (
	"Golang_Fiber/database"
	"Golang_Fiber/database/repositories"
	"Golang_Fiber/messages"
	"Golang_Fiber/utilities"
	"errors"
)

func GetOne(data interface{}, id uint, joins ...string) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	statement := database.DB.Model(data)
	statement = repositories.Joins(statement, joins)
	return statement.Limit(1).Find(data, id).Error
}

func GetFirst(data interface{}, joins ...string) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	statement := database.DB.Model(data)
	statement = repositories.Joins(statement, joins)
	return statement.First(data).Error
}

func GetOneByColumn(data interface{}, column string, value interface{}, joins ...string) error {
	if utilities.IsStruct(data) {
		return errors.New(messages.CrudDataNotPointer)
	}
	statement := database.DB.Model(data)
	statement = repositories.Joins(statement, joins)
	return statement.Limit(1).Find(data, column, value).Error
}
