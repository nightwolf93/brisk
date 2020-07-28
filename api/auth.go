package api

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/auth"
	"github.com/nightwolf93/brisk/storage"
)

type saveNewPairBody struct {
	ClientID     string `json:"client_id" xml:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" xml:"client_secret" form:"client_secret"`
}

// SaveNewPair save a new pair of credentials
func SaveNewPair(c *fiber.Ctx) {
	body := new(saveNewPairBody)
	if err := c.BodyParser(body); err != nil {
		c.SendStatus(400)
		return
	}
	if body.ClientID == "" || body.ClientSecret == "" {
		c.SendStatus(400)
		return
	}
	pair := &storage.ClientPairCredentials{
		ClientID:     body.ClientID,
		ClientSecret: body.ClientSecret,
	}

	// Check if is the master credential
	if !auth.IsMasterCredentials(c.Locals("credential").(*storage.ClientPairCredentials)) {
		c.SendStatus(401)
		return
	}

	// Check if there is already a pair for this client id
	if storage.FindSecretByID(pair.ClientID) != "" {
		c.SendStatus(409)
		return
	}

	// Save the pair
	storage.SavePair(pair)
	log.Printf("saved new credential clientId=%s clientSecret=%s", pair.ClientID, pair.ClientSecret)
}
