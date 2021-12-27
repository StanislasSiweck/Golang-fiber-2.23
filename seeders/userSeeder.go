package seeders

import (
	"Golang_Fiber/database"
	"Golang_Fiber/model"
)

func userSeeder() {
	user := []model.User{
		{
			Firstname: "Stanislas",
			Lastname:  "Siweck",
			Email:     "s.siweck@gmail.com",
		},
	}
	database.DB.Create(&user)
}
