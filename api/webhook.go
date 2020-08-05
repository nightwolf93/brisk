package api

import (
	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/storage"
	"github.com/nightwolf93/brisk/webhook"
)

type createWebhookBody struct {
	URL      string   `json:"url" xml:"url" form:"url"`
	Bindings []string `json:"bindings" xml:"bindings" form:"bindings"`
}

// RegisterWebhook register a new webhook
func RegisterWebhook(c *fiber.Ctx) {
	// Parse body
	body := new(createWebhookBody)
	if err := c.BodyParser(body); err != nil {
		c.SendStatus(400)
		return
	}
	w := &webhook.Webhook{
		URL:      body.URL,
		Bindings: body.Bindings,
		Owner:    c.Locals("credential").(*storage.ClientPairCredentials).ClientID,
	}
	webhook.AddWebhook(w)
}
