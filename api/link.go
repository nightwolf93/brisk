package api

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/nightwolf93/brisk/webhook"

	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/storage"
	"github.com/nightwolf93/brisk/utils"
)

type createLinkBody struct {
	URL        string   `json:"url" xml:"url" form:"url"`
	TTL        int      `json:"ttl" xml:"ttl" form:"ttl"`
	Slug       string   `json:"slug" xml:"slug" form:"slug"`
	SlugLength int      `json:"slug_length" xml:"slug_length" form:"slug_length"`
	Services   []string `json:"services" xml:"services" form:"services"`
}

type deleteLinkBody struct {
	Slug string `json:"slug" xml:"slug" form:"slug"`
}

type updateLinkBody struct {
	Slug string `json:"slug" xml:"slug" form:"slug"`
	URL  string `json:"url" xml:"url" form:"url"`
	TTL  int    `json:"ttl" xml:"ttl" form:"ttl"`
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

	// Check if this is a url
	if !isValidUrl(body.URL) {
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
		Slug:      slug,
		URL:       body.URL,
		TTL:       body.TTL,
		Owner:     c.Locals("credential").(*storage.ClientPairCredentials).ClientID,
		UpdatedAt: int32(time.Now().Unix()),
		CreateAt:  int32(time.Now().Unix()),
	}
	storage.SaveLink(link)

	// Call webhooks
	go webhook.CallWebhooks("new_link", map[string]interface{}{
		"slug":       slug,
		"owner":      c.Locals("credential").(*storage.ClientPairCredentials).ClientID,
		"updated_at": link.UpdatedAt,
		"created_at": link.CreateAt,
		"ttl":        link.TTL,
	})

	log.Printf("new link created slug=%s url=%s owner=%s", link.Slug, link.URL, link.Owner)
	c.JSON(map[string]interface{}{
		"slug": slug,
		"url":  fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), slug),
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

	visitorEntry := storage.GetVisitorEntryByFiberCtx(c)
	storage.SaveVisitorEntry(link, visitorEntry)
	log.Printf("visitor on link=%s visitor=%s ip=%s", link.Slug, visitorEntry.Hash, visitorEntry.IP)

	// Call webhooks
	go webhook.CallWebhooks("visit_link", map[string]interface{}{
		"link": map[string]interface{}{
			"slug":           link.Slug,
			"visitor_amount": link.VisitAmount,
			"url":            link.URL,
			"owner":          link.Owner,
		},
		"visitor": visitorEntry,
	})

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

func UpdateLink(c *fiber.Ctx) {
	// Parse body
	body := new(updateLinkBody)
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

	// Update link
	if body.TTL > 0 {
		link.TTL = body.TTL
	}
	link.URL = body.URL
	link.UpdatedAt = int32(time.Now().Unix())
	storage.SaveLink(link)

	// Call webhooks
	go webhook.CallWebhooks("update_link", map[string]interface{}{
		"slug":       link.Slug,
		"owner":      c.Locals("credential").(*storage.ClientPairCredentials).ClientID,
		"updated_at": link.UpdatedAt,
		"created_at": link.CreateAt,
		"ttl":        link.TTL,
	})

	log.Printf("link updated slug=%s", link.Slug)
	c.SendStatus(200)
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
