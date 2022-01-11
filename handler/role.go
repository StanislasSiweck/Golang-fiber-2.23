package handler

import (
	"Golang_Fiber/database/repositories/CRUD"
	"Golang_Fiber/messages"
	"Golang_Fiber/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

//GetAllRole Récupérer tous les roles
func GetAllRole(c *fiber.Ctx) error {
	defer recovery()
	var roles []model.Role
	joins := getJoins(c)

	if err := CRUD.GetAll(&roles, joins...); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le role avec l'ID", err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.Success, "data": roles})
}

//GetOneRole Récupére un role
func GetOneRole(c *fiber.Ctx) error {
	defer recovery()
	joins := getJoins(c)
	roleId, _ := c.ParamsInt("id")

	var role model.Role
	if err := CRUD.GetOne(&role, uint(roleId), joins...); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le role avec l'ID", err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.Success, "data": role})
}

//CreateRole Crée un role
func CreateRole(c *fiber.Ctx) error {
	defer recovery()

	var role model.Role
	if err := c.BodyParser(&role); err != nil {
		return SendBodyParseError(err)
	}
	if err := Validate.Struct(role); err != nil {
		return SendValidatorError(err)
	}

	if err := CRUD.Create(&role); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le role avec l'ID", err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.SuccessCreate, "data": role})
}

//UpdateRole Mettre à jour un role
func UpdateRole(c *fiber.Ctx) error {
	defer recovery()
	roleId, _ := c.ParamsInt("id")

	var roleUpdate model.Role
	if err := c.BodyParser(&roleUpdate); err != nil {
		return SendBodyParseError(err)
	}
	if err := Validate.Struct(roleUpdate); err != nil {
		return SendValidatorError(err)
	}

	var role model.Role
	if err := CRUD.GetOne(&role, uint(roleId)); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le role avec l'ID", err)
	}

	roleUpdate.Id = role.Id
	if err := CRUD.Update(&roleUpdate); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le role avec l'ID", err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.SuccessUpdate, "details": roleUpdate})
}

//DeleteRole Supprime un role
func DeleteRole(c *fiber.Ctx) error {
	defer recovery()
	roleId, _ := c.ParamsInt("id")

	var role model.Role
	if err := CRUD.GetOne(&role, uint(roleId)); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Impossible de trouver le role avec l'ID", err)
	}

	if err := CRUD.Delete(&role); err != nil {
		return SendError(fiber.StatusInternalServerError, "messages.DeleteContractError", "", err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": messages.SuccessDelete, "details": ""})
}
