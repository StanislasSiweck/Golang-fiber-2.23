package router

import (
	"Golang_Fiber/database"
	"Golang_Fiber/handler"
	"Golang_Fiber/middleware"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"time"
)

//routeGeneralConfig Config de route
func routeGeneralConfig(app *fiber.App) {
	//Panic recovery
	app.Use(recover.New())
	//Logger
	app.Use(logger.New())

	//Log de chaque requette
	app.Use(
		logger.New(logger.Config{
			Format:       "${cyan}[${time}] ${black}| ${yellow}${status} ${black}| ${blue}[${method}] ${black}${path} ${black}| ${red}${latency}\n",
			TimeFormat:   "02-Jan-2006 15:04:05",
			TimeZone:     "Europe/Paris",
			TimeInterval: 500 * time.Millisecond,
		}),
	)

	//Limite de requete
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			if c.Method() == "OPTIONS" {
				return true
			}
			return c.IP() == "127.0.0.1"
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		Max:        5,
		Expiration: 60 * time.Second,
	}))
}

//SetupRoutes Liste de routes
func SetupRoutes(app *fiber.App) {
	routeGeneralConfig(app)

	e := InitCasbin()
	_ = e.LoadPolicy()

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})

	api.Get("/login", handler.Login, middleware.CasbinMiddleware(e.Enforcer))

	user := api.Group("/user", middleware.ValidateAuth(), middleware.CasbinMiddleware(e.Enforcer))
	user.Get("/", handler.GetAllUser)
	user.Get("/current", handler.CurrentUser)
	user.Get("/:id", handler.GetOneUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	role := api.Group("/role", middleware.ValidateAuth())
	role.Get("/", handler.GetAllRole)
	role.Get("/:id", handler.GetOneRole)
	role.Post("/", handler.CreateRole)
	role.Put("/:id", handler.UpdateRole)
	role.Delete("/:id", handler.DeleteRole)
}

func InitCasbin() *casbin.CachedEnforcer {
	a, _ := gormadapter.NewAdapterByDB(database.DB)
	var err error
	var e *casbin.CachedEnforcer
	e, err = casbin.NewCachedEnforcer("./config/model.conf", a)
	if err != nil {
		log.Panic("Error loading enforcer")
	}

	_, _ = e.AddPolicy("user", "/api/*", "(GET)|(POST)|(PUT)|(DELETE)|(PATCH)")

	err = e.SavePolicy()
	if err != nil {
		log.Panic(err.Error())
	}
	return e
}
