package webhook

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Webhook represent a webhook
type Webhook struct {
	URL      string
	Bindings []string
}

var webhooks []*Webhook

// Init init webhook system
func Init() {
	webhooks = make([]*Webhook, 0)
}

// AddWebhook add a new webhook
func AddWebhook(webhook *Webhook) {
	webhooks = append(webhooks, webhook)
	log.Printf("new webhook created url=%s", webhook.URL)
}

// GetWebhooksByBinding return all webhooks with the binding required
func GetWebhooksByBinding(binding string) []*Webhook {
	hooks := make([]*Webhook, 0)
	for _, w := range webhooks {
		if w.hasBinding(binding) {
			hooks = append(hooks, w)
		}
	}
	return hooks
}

// CallWebhooks call webhooks with the given binding
func CallWebhooks(binding string, payload interface{}) {
	for _, webhook := range GetWebhooksByBinding(binding) {
		jsonStr, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", webhook.URL, bytes.NewBuffer(jsonStr))
		if err != nil {
			log.Printf("error when calling the webhook url=%s", webhook.URL)
			continue
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("error when calling the webhook url=%s error=%s", webhook.URL, err.Error())
			continue
		}
		defer resp.Body.Close()

		log.Printf("Status code %d", resp.Status)
	}
}

func (webhook *Webhook) hasBinding(binding string) bool {
	for _, v := range webhook.Bindings {
		if v == binding {
			return true
		}
	}
	return false
}
