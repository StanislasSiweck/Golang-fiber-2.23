package seeders

import (
	"Golang_Fiber/database/repositories/CRUD"
	"Golang_Fiber/model"
	"fmt"
)

func userRolesSeeder() {
	var users []model.User
	_ = CRUD.GetAll(&users)
	var role model.Role
	_ = CRUD.GetFirst(&role)

	for _, user := range users {
		user.RoleId = 1
		fmt.Println(user)
		//_ = CRUD.Update(&user)
	}
}
