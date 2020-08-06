package addon

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type shortestResponse struct {
	Status       string `json:"status"`
	ShortenedURL string `json:"shortenedUrl"`
}

// ConvertToShortest convert a brisk link to a shortest link
func ConvertToShortest(link string) string {
	form := url.Values{}
	form.Add("urlToShorten", link)
	req, err := http.NewRequest("PUT", "https://api.shorte.st/v1/data/url", strings.NewReader(form.Encode()))
	if err != nil {
		return ""
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("public-api-token", os.Getenv("SHORTEST_PUBLIC_TOKEN"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	payload := &shortestResponse{}
	json.NewDecoder(resp.Body).Decode(payload)
	return payload.ShortenedURL
}
