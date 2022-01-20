package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"lisxAPI/controllers"
	"lisxAPI/db"
	"lisxAPI/middlewares"
	"log"
	"os"
)

var resources []string

func registerResource(resource string) {
	resources = append(resources, resource)
}

func noEnvArg() bool {
	for _, arg := range os.Args {
		if arg == "--no-env" {
			return true
		}
	}
	return false
}

func main() {
	if !noEnvArg() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
	db.DB = db.Connect()
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/resources", middlewares.RequireAdmin, func(ctx *fiber.Ctx) error {
		return ctx.JSON(resources)
	})

	app.Post("/access-tokens", controllers.CreateAccessToken)

	app.Get("/api-keys", middlewares.RequireAdmin, controllers.GetAPIKeys)
	app.Post("/api-keys", middlewares.RequireAdmin, controllers.CreateAPIKey)

	app.Get("/api-key-permissions", middlewares.RequireAdmin, controllers.GetAPIKeyPermissions)
	app.Put("/api-key-permissions", middlewares.RequireAdmin, controllers.UpdateAPIKeyPermissions)

	app.Get("/analyzers", middlewares.RequireReadPermission, controllers.GetAnalyzers)
	app.Get("/analyzers/:id", middlewares.RequireReadPermission, controllers.GetAnalyzerById)
	app.Post("/analyzers", middlewares.RequireCreatePermission, controllers.CreateAnalyzer)

	app.Get("/users", middlewares.RequireReadPermission, controllers.GetUsers)
	app.Get("/users/:id", middlewares.RequireReadPermission, controllers.GetUserById)
	app.Post("/users", middlewares.RequireAdmin, controllers.CreateUser)

	app.Get("/user-permissions", middlewares.RequireAdmin, controllers.GetUserPermissions)
	app.Put("/user-permissions", middlewares.RequireAdmin, controllers.UpdateUserPermissions)

	registerResource("/analyzers")
	registerResource("/users")
	log.Fatal(app.Listen(":8080"))
}
