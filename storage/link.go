package storage

type Link struct {
	Slug              string
	URL               string
	Owner             string
	TTL               int
	CreateAtTimestamp int32
}
