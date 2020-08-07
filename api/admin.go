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

// AdminGetWebhooks get all active webhooks
func AdminGetWebhooks(c *fiber.Ctx) {
	c.JSON(webhook.GetWebhooks())
}

// AdminGetVisitors get all visitors for a link
func AdminGetVisitors(c *fiber.Ctx) {
	link := storage.FindLink(c.Params("slug"))
	if link == nil { // Link not found
		c.SendStatus(400)
		return
	}

	visitors, err := storage.FindVisitorsForLink(link)
	if err != nil {
		c.SendStatus(500)
		return
	}
	c.JSON(visitors)
}
