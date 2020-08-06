package api

import (
	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/storage"
	"github.com/nightwolf93/brisk/webhook"
)

// AdminGetAllLinks get all links stored into the db
func AdminGetAllLinks(c *fiber.Ctx) {
	links, err := storage.FindAllLinks()
	if err != nil {
		c.SendStatus(500)
		return
	}
	c.JSON(links)
}

func AdminGetWebhooks(c *fiber.Ctx) {
	c.JSON(webhook.GetWebhooks())
}
