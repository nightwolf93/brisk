package storage

type Link struct {
	Slug              string `json:"slug"`
	URL               string `json:"url"`
	Owner             string `json:"owner"`
	TTL               int    `json:"ttl"`
	CreateAtTimestamp int32  `json:"create_at_timestamp"`
}
