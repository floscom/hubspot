package hubspot

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	APIKey     string     `json:"-"`
	Email      string     `json:"-"`
	Properties []Property `json:"properties"`
}

type ContactResp struct {
	Vid   int  `json:"vid"`
	IsNew bool `json:"isNew"`
}

func (h *Contact) Add(prop, value string) {
	h.Properties = append(h.Properties, Property{prop, value})
}

func NewContact(apiKey, email string) *Contact {
	return &Contact{
		APIKey: apiKey,
		Email:  email,
	}
}

func (h *Contact) Publish() (cr *ContactResp) {
	url := fmt.Sprintf("/contacts/v1/contact/createOrUpdate/email/%s/?hapikey=%s", h.Email, h.APIKey)

	b, _ := json.Marshal(h)

	x := Send(url, "POST", b)

	cr = &ContactResp{}
	err := json.Unmarshal(x, cr)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return
}
