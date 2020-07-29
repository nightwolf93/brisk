package api

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/storage"
	"github.com/nightwolf93/brisk/utils"
)

type createLinkBody struct {
	URL string `json:"url" xml:"url" form:"url"`
	TTL int    `json:"ttl" xml:"ttl" form:"ttl"`
}

// CreateLink request to create a new link
func CreateLink(c *fiber.Ctx) {
	// Parse body
	body := new(createLinkBody)
	if err := c.BodyParser(body); err != nil {
		c.SendStatus(400)
		return
	}

	// Check the TTL requested
	maxTTL, _ := strconv.Atoi(os.Getenv("MAX_LINK_TTL"))
	if body.TTL > maxTTL {
		c.SendStatus(400)
		return
	}

	// Create the link
	slug := utils.RandomString(7)
	link := &storage.Link{
		Slug:              slug,
		URL:               body.URL,
		TTL:               body.TTL,
		Owner:             c.Locals("credential").(*storage.ClientPairCredentials).ClientID,
		CreateAtTimestamp: int32(time.Now().Unix()),
	}
	storage.SaveLink(link)
	log.Printf("new link created slug=%s url=%s owner=%s", link.Slug, link.URL, link.Owner)
	c.JSON(map[string]interface{}{
		"slug": slug,
	})
}

func GetLink(c *fiber.Ctx) {
	link := storage.FindLink(c.Params("slug"))
	if link == nil {
		c.Send("Not found")
		return
	}
	c.Redirect(link.URL)
}
