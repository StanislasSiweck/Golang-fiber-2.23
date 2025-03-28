package handler

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

//SendBodyParseError
func SendBodyParseError(error error) error {
	if os.Getenv("ENV") == "LOCAL" {
		return SendError(fiber.StatusBadRequest, "messages.ErrorDecoding", error.Error(), error)
	}
	return SendError(fiber.StatusBadRequest, "messages.ErrorDecoding", "messages.CheckTheLogs", error)
}

//SendValidatorError Returns a validator error
func SendValidatorError(error error) error {
	return SendError(fiber.StatusUnprocessableEntity, "messages.ErrorValidating", error.Error(), error)
}

func getJoins(c *fiber.Ctx) []string {
	join := c.Query("with")
	return strings.Split(join, ",")
}
