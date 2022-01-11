package handler

import (
	"Golang_Fiber/database/repositories/CRUD"
	"Golang_Fiber/messages"
	"Golang_Fiber/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

//GetAllUser Récupérer tous les utilisateurs
func GetAllUser(c *fiber.Ctx) error {
	defer recovery()
	var users []model.User
	joins := getJoins(c)

	if err := CRUD.GetAll(&users, joins...); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le user avec l'ID", err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.Success, "data": users})
}

//GetOneUser Récupére un utilisateur
func GetOneUser(c *fiber.Ctx) error {
	defer recovery()
	joins := getJoins(c)
	userId, _ := c.ParamsInt("id")

	var user model.User
	if err := CRUD.GetOne(&user, uint(userId), joins...); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le user avec l'ID", err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.Success, "data": user})
}

//CreateUser Crée un utilisateur
func CreateUser(c *fiber.Ctx) error {
	defer recovery()

	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return SendBodyParseError(err)
	}
	if err := Validate.Struct(user); err != nil {
		return SendValidatorError(err)
	}

	if err := CRUD.Create(&user); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le user avec l'ID", err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.SuccessCreate, "data": user})
}

//UpdateUser Mettre à jour un utilisateur
func UpdateUser(c *fiber.Ctx) error {
	defer recovery()
	userId, _ := c.ParamsInt("id")

	var userUpdate model.User
	if err := c.BodyParser(&userUpdate); err != nil {
		return SendBodyParseError(err)
	}
	if err := Validate.Struct(userUpdate); err != nil {
		return SendValidatorError(err)
	}

	var user model.User
	if err := CRUD.GetOne(&user, uint(userId)); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le user avec l'ID", err)
	}

	userUpdate.Id = user.Id
	if err := CRUD.Update(&userUpdate); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le user avec l'ID", err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.SuccessUpdate, "details": userUpdate})
}

//DeleteUser Supprime un utilisateur
func DeleteUser(c *fiber.Ctx) error {
	defer recovery()
	userId, _ := c.ParamsInt("id")

	var user model.User
	if err := CRUD.GetOne(&user, uint(userId)); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le user avec l'ID", err)
	}

	if err := CRUD.Delete(&user); err != nil {
		return SendError(fiber.StatusInternalServerError, "messages.DeleteContractError", "", err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.SuccessDelete, "details": ""})
}

func CurrentUser(c *fiber.Ctx) error {
	joins := getJoins(c)
	user := getCurrentUser(c, joins...)

	if user.Id != 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.Success, "data": user})
	}
	return SendError(fiber.StatusNotFound, messages.Error, "Utilisateur pas trouver")
}
