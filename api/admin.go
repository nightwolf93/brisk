package api

import (
	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/storage"
)

func AdminGetAllLinks(c *fiber.Ctx) {
	links, err := storage.FindAllLinks()
	if err != nil {
		c.SendStatus(500)
		return
	}
	c.JSON(links)
}
