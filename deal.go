package hubspot

import (
	"encoding/json"
	"fmt"
	"log"
)

type DealResp struct {
	Vid int `json:"dealId"`
}

type Deal struct {
	APIKey string `json:"-"`

	Associations struct {
		AssociatedCompanyIds []int `json:"associatedCompanyIds,omitempty"`
		AssociatedVids       []int `json:"associatedVids,omitempty"`
	} `json:"associations,omitempty"`

	PortalID int `json:"portalId,omitempty"`

	Properties []PropertyDeal `json:"properties,omitempty"`
}

func NewDeal(apiKey string) *Deal {
	return &Deal{
		APIKey: apiKey,
	}
}

func (h *Deal) Add(prop string, value interface{}) {
	h.Properties = append(h.Properties, PropertyDeal{prop, value})
}

func (h *Deal) Publish() (dr *DealResp) {
	url := fmt.Sprintf("/deals/v1/deal?hapikey=%s", h.APIKey)

	b, _ := json.Marshal(h)

	x := Send(url, "POST", b)

	log.Println(string(b))

	dr = &DealResp{}
	err := json.Unmarshal(x, dr)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return
}
