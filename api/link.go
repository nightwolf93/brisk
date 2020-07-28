package api

import (
	"log"
	"time"

	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/storage"
	"github.com/nightwolf93/brisk/utils"
)

type createLinkBody struct {
	URL string `json:"url" xml:"url" form:"url"`
}

// CreateLink request to create a new link
func CreateLink(c *fiber.Ctx) {
	// Parse body
	body := new(createLinkBody)
	if err := c.BodyParser(body); err != nil {
		c.SendStatus(400)
		return
	}

	// Create the link
	slug := utils.RandomString(7)
	link := &storage.Link{
		Slug:             slug,
		URL:              body.URL,
		CreateAtTimetamp: int32(time.Now().Unix()),
	}
	log.Printf("new link created slug=%s url=%s", link.Slug, link.URL)
	c.JSON(map[string]interface{}{
		"slug": slug,
	})
}
