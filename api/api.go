package api

import (
	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/auth/middleware"
)

// Init the api http app
func Init() {
	app := fiber.New(&fiber.Settings{
		DisableStartupMessage: true,
	})

	// Apply middleware
	app.Use("/api", middleware.CheckCredential)

	app.Put("/api/v1/credential", SaveNewPair)
	app.Put("/api/v1/link", CreateLink)
	app.Delete("/api/v1/link", DeleteLink)
	app.Get("/:slug", GetLink)

	app.Listen(3000)
}
