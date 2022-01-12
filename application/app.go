package application

import (
	"Golang_Fiber/database"
	"Golang_Fiber/handler"
	"Golang_Fiber/router"
	"Golang_Fiber/seeders"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

func InitFiberApp() *fiber.App {
	app := fiber.New(
		fiber.Config{
			BodyLimit: 5000 * 1024 * 1024,
			AppName:   "BASE API",
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {

				code := fiber.StatusInternalServerError
				content := make([]string, 0)
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
					content = strings.Split(e.Message, ";")
				}

				if len(content) > 2 {
					log.Println(content[2])
				} else if len(content) < 2 && len(content) > 0 {
					log.Println(content)
					return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": content,
					})
				} else if len(content) == 0 {
					return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "loggingError",
					})
				}
				return ctx.Status(code).JSON(fiber.Map{
					"message": content[0],
					"details": content[1],
				})
			},
		},
	)

	database.ConnectDB()
	handler.InitValidator()
	router.SetupRoutes(app)

	return app
}

//CliCommandApp	Migration et Seeder de la BDD
func CliCommandApp(fresh bool, seed bool) {
	//Se connecter à la base de données
	database.ConnectDB()

	//Migration
	if fresh {
		database.DropTables()
		database.MigrateDatabase()
	}
	//Seeder
	if seed {
		seeders.Seed()
	}
}
