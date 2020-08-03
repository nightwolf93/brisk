package api

import (
	"github.com/nightwolf93/brisk/webhook"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/storage"
	"github.com/nightwolf93/brisk/utils"
)

type createLinkBody struct {
	URL        string `json:"url" xml:"url" form:"url"`
	TTL        int    `json:"ttl" xml:"ttl" form:"ttl"`
	Slug       string `json:"slug" xml:"slug" form:"slug"`
	SlugLength int    `json:"slug_length" xml:"slug_length" form:"slug_length"`
}

type deleteLinkBody struct {
	Slug string `json:"slug" xml:"slug" form:"slug"`
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

	// Check slug length
	if body.SlugLength > 20 || body.SlugLength < 3 {
		c.SendStatus(400)
		return
	}

	// Create the link
	slug := ""
	if len(body.Slug) == 0 {
		for i := 0; i < 32; i++ {
			slug = utils.RandomString(body.SlugLength)
			if storage.FindLink(slug) == nil {
				break
			}
		}
	} else {
		if len(body.Slug) > 20 {
			c.SendStatus(400)
			return
		}
		slug = body.Slug
	}
	if storage.FindLink(slug) != nil {
		c.SendStatus(409)
		return
	}

	link := &storage.Link{
		Slug:              slug,
		URL:               body.URL,
		TTL:               body.TTL,
		Owner:             c.Locals("credential").(*storage.ClientPairCredentials).ClientID,
		CreateAtTimestamp: int32(time.Now().Unix()),
	}
	storage.SaveLink(link)

	// Call webhooks
	webhook.CallWebhooks("new_link", map[string]interface{}{
		"slug": slug,
	})

	log.Printf("new link created slug=%s url=%s owner=%s", link.Slug, link.URL, link.Owner)
	c.JSON(map[string]interface{}{
		"slug": slug,
	})
}

// GetLink redirect to a link by the slug
func GetLink(c *fiber.Ctx) {
	link := storage.FindLink(c.Params("slug"))
	if link == nil {
		c.Send("Not found")
		return
	}
	link.VisitAmount = link.VisitAmount + 1
	storage.SaveLink(link)
	c.Redirect(link.URL)
}

// DeleteLink request to delete a link
func DeleteLink(c *fiber.Ctx) {
	// Parse body
	body := new(deleteLinkBody)
	if err := c.BodyParser(body); err != nil {
		c.SendStatus(400)
		return
	}

	link := storage.FindLink(body.Slug)

	// Check if the link exist
	if link == nil {
		c.SendStatus(400)
		return
	}

	// Check if is the link owner
	if link.Owner != c.Locals("credential").(*storage.ClientPairCredentials).ClientID {
		c.SendStatus(400)
		return
	}

	log.Printf("link deleted slug=%s", link.Slug)
	storage.DeleteLink(link.Slug)

	c.SendStatus(200)
}
