package seeders

import (
	"Golang_Fiber/database"
	"Golang_Fiber/database/repositories/CRUD"
	"Golang_Fiber/model"
	"Golang_Fiber/utilities"
)

func userSeeder() {
	var role model.Role
	_ = CRUD.GetFirst(&role)

	user := []model.User{
		{
			Firstname: "Stanislas",
			Lastname:  "Siweck",
			Email:     "s.siweck@gmail.com",
			Password:  utilities.HashString("s.siweck@gmail.com"),
			RoleId:    role.Id,
		},
	}
	database.DB.Create(&user)
}
