package middleware

import (
	"Golang_Fiber/jwtauth"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"strings"
)

type CasbinConfig struct {
	Enforcer *casbin.Enforcer
}

/*
	CasbinMiddleware:
	This middleware is HEAVILY inspired by:  https://github.com/prongbang/fiber-casbinrest
	credits to the developers
*/
func CasbinMiddleware(e *casbin.Enforcer) fiber.Handler {
	config := CasbinConfig{
		Enforcer: e,
	}
	return func(c *fiber.Ctx) error {
		if c.Get("Authorization") == os.Getenv("INTERNAL_AKT") || config.CheckPermissions(c) {
			return c.Next()
		}
		if os.Getenv("APP_ENV") == "LOCAL" {
			log.Println("Error: CasbinMiddleware")
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden."})
	}
}

/*
	GetRoleByToken:
	Loops through the JWT role and checks the user role for the app
*/
func (config *CasbinConfig) GetRoleByToken(c *fiber.Ctx) string {
	role := "anonymous"
	jwt := jwtauth.GetLoggedUserData(c)
	if jwt.Role != "" {
		role = jwt.Role
	}
	return role
}

/*
	CheckPermissions:
	Check if casbin allows the user or not to access the method
*/
func (config *CasbinConfig) CheckPermissions(c *fiber.Ctx) bool {
	role := config.GetRoleByToken(c)
	allowed := false

	result, err := config.Enforcer.Enforce(strings.ToLower(role), c.Path(), c.Method())
	if result && err == nil {
		allowed = true
	} else {
		log.Println(err)
	}
	return allowed
}
