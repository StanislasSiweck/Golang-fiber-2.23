package seeders

import (
	"Golang_Fiber/database"
	"Golang_Fiber/model"
)

func roleSeeder() {
	roles := []model.Role{
		{
			Label: "User",
		},
	}
	database.DB.Create(&roles)
}
