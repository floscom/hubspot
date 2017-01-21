package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
  "log"
)

type SingleSendEmail struct {
	APIKey     string             `json:"-"`
	EmailId    int                `json:"emailId"`
  Message    Message            `json:"message"`
	ContactProperties []Property  `json:"contactProperties"`
  CustomProperties  []Property  `json:"customProperties"`
}

type Message struct {
  To         string             `json:"to"`
}

type SingleSendEmailResponse struct {
	SendResult   string  `json:"send_result"`
	Message      string  `json:"message"`
  Id           int     `json:"id"`
}

func NewEmail(api_key string, email_id int, to string) *SingleSendEmail {
	return &SingleSendEmail{
		APIKey:   api_key,
		EmailId:  email_id,
    Message:  Message{To: to},
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
	url := fmt.Sprintf("https://api.hubapi.com/email/public/v1/singleEmail/send?hapikey=%s", h.APIKey)

	b, _ := json.Marshal(h)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HUBSPOT ERROR: ", err)
	}
	defer resp.Body.Close()

	x, _ := ioutil.ReadAll(resp.Body)

	ssr = &SingleSendEmailResponse{}
	err = json.Unmarshal(x, ssr)
	if err != nil {
		log.Println("HUBSPOT ERROR: ", err)
		return nil
	}
	return ssr
}