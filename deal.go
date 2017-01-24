package hubspot

import (
	"encoding/json"
	"fmt"
)

type Deal struct {
	APIKey string `json:"-"`

	Associations struct {
		AssociatedCompanyIds []int `json:"associatedCompanyIds,omitempty"`
		AssociatedVids       []int `json:"associatedVids,omitempty"`
	} `json:"associations,omitempty"`

	PortalID int `json:"portalId,omitempty"`

	Properties []Property `json:"properties,omitempty"`
}

func NewDeal(apiKey string) *Deal {
	return &Deal{
		APIKey: apiKey,
	}
}

func (h *Deal) Add(prop string, value interface{}) {
	h.Properties = append(h.Properties, Property{prop, value})
}

func (h *Deal) Publish() {
	url := fmt.Sprintf("/deals/v1/deal?hapikey=%s", h.APIKey)

	b, _ := json.Marshal(h)

	x := Send(url, "POST", b)

	fmt.Println("Hubspot body", string(x))
}
