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
	app.Use("/api/v1/admin", middleware.CheckMasterCredential)

	app.Put("/api/v1/credential", SaveNewPair)
	app.Put("/api/v1/link", CreateLink)
	app.Delete("/api/v1/link", DeleteLink)
	app.Get("/:slug", GetLink)

	app.Get("/api/v1/admin/link", AdminGetAllLinks)

	app.Listen(3000)
}
