package hubspot

import (
	"encoding/json"
	"fmt"
)

type SingleSendEmail struct {
	APIKey            string     `json:"-"`
	EmailId           int        `json:"emailId"`
	Message           Message    `json:"message"`
	ContactProperties []Property `json:"contactProperties"`
	CustomProperties  []Property `json:"customProperties"`
}

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	ReplyTo string `json:"reply_to"`
}

type SingleSendEmailResponse struct {
	SendResult string `json:"send_result"`
	Message    string `json:"message"`
	Id         int    `json:"id"`
}

func NewEmail(api_key string, email_id int, message Message) *SingleSendEmail {
	return &SingleSendEmail{
		APIKey:  api_key,
		EmailId: email_id,
		Message: message,
	}
}

func (h *SingleSendEmail) Add(prop_type, prop, value string) {
	if prop_type == "contact" {
		h.ContactProperties = append(h.ContactProperties, Property{prop, value})
	}

	if prop_type == "custom" {
		h.CustomProperties = append(h.CustomProperties, Property{prop, value})
	}
}

func (h *SingleSendEmail) Publish() (ssr *SingleSendEmailResponse) {
	url := fmt.Sprintf("/email/public/v1/singleEmail/send?hapikey=%s", h.APIKey)

	b, _ := json.Marshal(h)

	x := Send(url, "POST", b)

	ssr = &SingleSendEmailResponse{}
	json.Unmarshal(x, ssr)

	return ssr
}
