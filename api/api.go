package api

import (
	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/auth/middleware"
	"log"
	"os"
	"strconv"
)

// Init the api http app
func Init() {
	app := fiber.New(&fiber.Settings{
		DisableStartupMessage: true,
	})

	// Apply middleware
	app.Use("/api", middleware.CheckCredential)
	app.Use("/api/v1/admin", middleware.CheckMasterCredential)

	// Link
	app.Put("/api/v1/link", CreateLink)
	app.Delete("/api/v1/link", DeleteLink)
	app.Get("/:slug", GetLink)

	// Credentials
	app.Put("/api/v1/credential", SaveNewPair)

	// Webhook
	app.Put("/api/v1/webhook", RegisterWebhook)

	// Admin
	app.Get("/api/v1/admin/link", AdminGetAllLinks)

	port := 3000
	if os.Getenv("PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	log.Printf("listen on *:%d", port)
	app.Listen(port)
}
