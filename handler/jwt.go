package handler

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"strconv"
	"time"
)

//DefaultJWT Création de JWT
func DefaultJWT(c *fiber.Ctx) error {
	//type MapClaims map[string]interface{}
	ttlJwt, _ := strconv.Atoi(os.Getenv("JWT_TTL"))

	atClaims := jwt.MapClaims{}
	atClaims["exp"] = time.Now().Add(time.Second * time.Duration(ttlJwt)).Unix()

	//Build JWT
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	//Signature du JWT
	accessToken, _ := at.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	respData := map[string]string{
		"access_token": accessToken,
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "connectionok", "data": respData})
}
