package seeders

import (
	"Golang_Fiber/database/repositories/CRUD"
	"Golang_Fiber/model"
)

func userRolesSeeder() {
	var users []model.User
	_ = CRUD.GetAll(&users)
	var roles []model.Role
	_ = CRUD.GetAll(&roles)

	for _, user := range users {
		user.Roles = roles
		_ = CRUD.Update(&user)
	}
}
