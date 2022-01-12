package middleware

import (
	"Golang_Fiber/jwtauth"
	"Golang_Fiber/messages"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
)

//ValidateAuth Check si le JWT est correct
func ValidateAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {

		key := c.Get("Authorization")
		_, err := jwt.Parse(key, jwtauth.JwtChecker())
		if err != nil {
			log.Print("Accès non autorisé à l'API, l'IP du demandeur est: " + c.IP() + ", la clé api utilisée était: " + key + ".")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": messages.InvalidAPIKey, "details": "Clé API invalide."})
		}
		return c.Next()
	}
}
