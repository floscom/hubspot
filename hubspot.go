package hubspot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Send(endpoint string, http_type string, data []byte) []byte {
	hubspot_url := fmt.Sprintf("https://api.hubapi.com%s", endpoint)
	req, err := http.NewRequest(http_type, hubspot_url, bytes.NewBuffer(data))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HUBSPOT ERROR: ", err)
	}
	defer resp.Body.Close()
	x, _ := ioutil.ReadAll(resp.Body)
	return x
}

func SendForm(endpoint string, http_type string, data url.Values) []byte {
	hubspot_url := fmt.Sprintf("https://forms.hubspot.com%s", endpoint)
	req, err := http.NewRequest(http_type, hubspot_url, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Accept", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HUBSPOT ERROR: ", err)
	}
	defer resp.Body.Close()
	x, _ := ioutil.ReadAll(resp.Body)
	return x
}
