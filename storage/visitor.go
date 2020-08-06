package storage

import (
	"github.com/alecthomas/geoip"
	"github.com/gofiber/fiber"
	"log"
	"net"
	"strings"
	"time"
)

// VisitorEntry is a entry for a unique visitor
type VisitorEntry struct {
	IP        string `json:"ip"`
	Referrer  string `json:"referrer"`
	Location  string `json:"location"`
	Timestamp int32  `json:"timestamp"`
}

// GetVisitorEntryByFiberCtx get a new visitor entry from a fiber context
func GetVisitorEntryByFiberCtx(c *fiber.Ctx) *VisitorEntry {
	ip := ""
	if string(c.Fasthttp.Request.Header.Peek("X-Forwarded-For")) != "" {
		ip = strings.Split(string(c.Fasthttp.Request.Header.Peek("X-Forwarded-For")), ",")[0]
	} else {
		ip = c.IP()
	}

	visitor := &VisitorEntry{}
	visitor.IP = ip
	visitor.Referrer = string(c.Fasthttp.Referer())
	visitor.Timestamp = int32(time.Now().Unix())

	// Check location
	if c.IP() != "127.0.0.1" {
		geo, err := geoip.New()
		if err == nil {
			country := geo.Lookup(net.ParseIP(ip))
			if country != nil {
				visitor.Location = country.Long
			} else {
				visitor.Location = "Not found"
			}
		} else {
			log.Printf("can't determine the location for the visitor: %s", err.Error())
		}
	} else {
		visitor.Location = "Localhost"
	}

	return visitor
}
