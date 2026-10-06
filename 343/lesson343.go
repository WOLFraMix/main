package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ExchangeRateResponse struct {
	Result             string             `json:"result"`
	Provider           string             `json:"provider"`
	Documentation      string             `json:"documentation"`
	TermsOfUse         string             `json:"terms_of_use"`
	TimeLastUpdateUnix int64              `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64              `json:"time_next_update_unix"`
	TimeNextUpdateUTC  string             `json:"time_next_update_utc"`
	TimeEOLUnix        int64              `json:"time_eol_unix"`
	BaseCode           string             `json:"base_code"`
	Rates              map[string]float64 `json:"rates"`
}

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://open.er-api.com/v6/latest/USD", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response ExchangeRateResponse
	if err := decoder.Decode(&response); err != nil {
		log.Fatal(err)
	}

	count := 0
	for r := range response.Rates {
		if r[0] == 'T' {
			count++
		}
	}
	fmt.Println(count)
}
