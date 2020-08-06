package storage

import (
	"github.com/alecthomas/geoip"
	"github.com/gofiber/fiber"
	"log"
	"net"
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
	visitor := &VisitorEntry{}
	visitor.IP = c.IP()
	visitor.Referrer = string(c.Fasthttp.Referer())
	visitor.Timestamp = int32(time.Now().Unix())

	// Check location
	if c.IP() != "127.0.0.1" {
		geo, err := geoip.New()
		if err == nil {
			country := geo.Lookup(net.ParseIP(c.IP()))
			visitor.Location = country.Long
		} else {
			log.Printf("can't determine the location for the visitor: %s", err.Error())
		}
	} else {
		visitor.Location = "Localhost"
	}
	log.Printf("%s", visitor.Location)

	return visitor
}
