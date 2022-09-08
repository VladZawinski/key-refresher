package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ListingResponse []struct {
	Coin Coin
}
type ListingBody struct {
	Currency string `json:"currency"`
	Sort     string `json:"sort"`
	Order    string `json:"order"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Meta     bool   `json:"meta"`
}

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
	request.Header.Add("x-api-key", "300842c5-aee8-4a55-8da9-efc75c6d8bcc")
	request.Header.Add("content-type", "application/json")
	if err != nil {
		fmt.Print(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	var coins []Coin
	responseBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBody, &coins)
	defer resp.Body.Close()

	return coins
}

func FakeParse() {
	b := []byte(`{"code": "BTC", "rate": 19212.045766750387}`)
	var coin Coin
	json.Unmarshal(b, &coin)
	log.Println(coin)

}
