package middleware

import (
	"github.com/gofiber/fiber"
	"github.com/nightwolf93/brisk/auth"
	"github.com/nightwolf93/brisk/storage"
)

// CheckMasterCredential is a middleware for cheking credential given by the client is a master
func CheckMasterCredential(c *fiber.Ctx) {
	clientID := string(c.Fasthttp.Request.Header.Peek("x-client-id"))
	clientSecret := string(c.Fasthttp.Request.Header.Peek("x-client-secret"))
	credential := &storage.ClientPairCredentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	isVerified := auth.VerifyPair(credential)
	if !isVerified {
		c.SendStatus(401)
		return
	}
	if !auth.IsMasterCredentials(credential) {
		c.SendStatus(401)
		return
	}
	c.Locals("credential", credential)
	c.Next()
}
