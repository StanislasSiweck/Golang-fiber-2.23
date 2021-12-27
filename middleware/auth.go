package middleware

import (
	"Golang_Fiber/messages"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

//ValidateAuth Check si le JWT est correct
func ValidateAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {

		key := c.Get("Authorization")
		_, err := jwt.Parse(key, jwtChecker())
		if err != nil {
			log.Print("Accès non autorisé à l'API, l'IP du demandeur est: " + c.IP() + ", la clé api utilisée était: " + key + ".")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": messages.InvalidAPIKey, "details": "Clé API invalide."})
		}
		return c.Next()
	}
}

//jwtChecker keyFunc will receive the parsed token and should return the key for validating.
func jwtChecker() func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	}
}
