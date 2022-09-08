package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Coin struct {
	Code string  `json:"code"`
	Rate float32 `json:"rate"`
}

func GetListing() []Coin {
	client := &http.Client{}
	payload := strings.NewReader(`{
		"currency": "USD",
		"sort": "rank",
		"order": "ascending",
		"offset": 0,
		"limit": 2,
		"meta": false
	}`)

	request, err := http.NewRequest(http.MethodPost, "https://api.livecoinwatch.com/coins/list/", payload)
	if err != nil {
		fmt.Print(err)
	}
	request.Header.Add("x-api-key", "300842c5-aee8-4a55-8da9-efc75c6d8bcc")
	request.Header.Add("content-type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var coins []Coin
	responseBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBody, &coins)

	return coins
}
