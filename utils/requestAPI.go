package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

func RequestGET(url string) string {

	base_url := os.Getenv("API_URL") + url

	// set json post if you want
	jsonPost := ""

	response, err := http.NewRequest("GET", base_url, bytes.NewBufferString(jsonPost))
	// to set api key if you want
	response.Header.Set("api-key", os.Getenv("API_KEY"))
	response.Header.Set("Content-Type", "application/json")

	if err != nil {
		return "upps error response"
	} else {
		client := &http.Client{}
		resp, err := client.Do(response)

		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		return string(data)
	}
}

func RequestPOST(id string) string {

	url := os.Getenv("API_URL")

	// set json post if you want
	jsonPost := `{
		"name": "example"
	}`

	response, err := http.NewRequest("POST", url, bytes.NewBufferString(jsonPost))
	// to set api key if you want
	response.Header.Set("api-key", os.Getenv("API_KEY"))
	response.Header.Set("Content-Type", "application/json")

	if err != nil {
		return "upps error response"
	} else {
		client := &http.Client{}
		resp, err := client.Do(response)

		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		return string(data)
	}
}
