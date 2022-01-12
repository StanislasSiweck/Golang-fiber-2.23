package jwtauth

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

type LoggedUser struct {
	Id   float64
	Role string
}

/*
	GetLoggedUserData:
	Gets the logged user data
*/
func GetLoggedUserData(c *fiber.Ctx) LoggedUser {
	var LoggedUser LoggedUser
	key := c.Get("Authorization")
	token, _ := jwt.Parse(key, JwtChecker())
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		id := claims["id"].(float64)
		role := claims["role"].(string)
		LoggedUser.Id = id
		LoggedUser.Role = role
		return LoggedUser
	} else if !token.Valid {
		log.Println("JWT invalide: ", token)
	}
	return LoggedUser
}

//JwtChecker keyFunc will receive the parsed token and should return the key for validating.
func JwtChecker() func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	}
}
