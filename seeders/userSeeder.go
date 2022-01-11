package seeders

import (
	"Golang_Fiber/database"
	"Golang_Fiber/model"
	"Golang_Fiber/utilities"
)

func userSeeder() {
	user := []model.User{
		{
			Firstname: "Stanislas",
			Lastname:  "Siweck",
			Email:     "s.siweck@gmail.com",
			Password:  utilities.HashString("s.siweck@gmail.com"),
		},
	}
	database.DB.Create(&user)
}
