package hubspot

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
	"log"
)

func Send(endpoint string, http_type string, data []byte) []byte {
	url := fmt.Sprintf("https://api.hubapi.com%s", endpoint)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
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