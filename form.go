package hubspot

import (
  "encoding/json"
  "fmt"
  "net/url"
)

type Form struct {
	APIKey     string     `json:"-"`
  PortalId   int        `json:"-"`
  FormGuid   string     `json:"-"`
	Properties []Property `json:"properties"`
}

type FormResp struct {
	HttpCode  int  `json:"http_code"`
}

func NewForm(api_key string, portal_id int, form_guid string) *Form {
	return &Form{
		APIKey:   api_key,
    PortalId: portal_id,
    FormGuid: form_guid,
	}
}

func (h *Form) Add(prop, value string) {
	h.Properties = append(h.Properties, Property{prop, value})
}

func (h *Form) Publish() (fr *FormResp) {
	hubspot_url := fmt.Sprintf("/uploads/form/v2/%d/%s?hapikey=%s", h.PortalId, h.FormGuid, h.APIKey)
  data := make(url.Values)
  for _, property := range h.Properties {
    data.Set(property.Property, property.Value.(string))
  }
	x := SendForm(hubspot_url, "POST", data)
	fr = &FormResp{}
	err := json.Unmarshal(x, fr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return
}