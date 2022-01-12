package handler

import (
	"Golang_Fiber/database/repositories/CRUD"
	"Golang_Fiber/jwtauth"
	"Golang_Fiber/messages"
	"Golang_Fiber/model"
	"Golang_Fiber/utilities"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Login et création de JWT
func Login(c *fiber.Ctx) error {
	var user model.User
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&login); err != nil {
		return SendBodyParseError(err)
	}

	if err := CRUD.GetOneByColumn(&user, "email", login.Email, "Role"); err != nil {
		return SendError(fiber.StatusNotFound, messages.Error, "Utilisateur na pas étais trouver", err)
	}

	//Validate password
	validPwd := utilities.CheckPassword([]byte(user.Password), login.Password)
	if !validPwd {
		return SendError(fiber.StatusNotFound, messages.Error, "Le mot de passe ou l'email est incorrecte")
	}
	//type MapClaims map[string]interface{}
	ttlJwt, _ := strconv.Atoi(os.Getenv("JWT_TTL"))

	atClaims := jwt.MapClaims{}
	atClaims["exp"] = time.Now().Add(time.Second * time.Duration(ttlJwt)).Unix()
	atClaims["id"] = user.Id
	atClaims["role"] = user.Role.Label

	//Build JWT
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	//Signature du JWT
	accessToken, _ := at.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	respData := map[string]string{
		"access_token": accessToken,
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "connectionok", "data": respData})
}

func getCurrentUser(c *fiber.Ctx, joins ...string) model.User {
	jwt := jwtauth.GetLoggedUserData(c)
	var user model.User
	if err := CRUD.GetOneByColumn(&user, "id", jwt.Id); err != nil {
		return user
	}
	return user
}
