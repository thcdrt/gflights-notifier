package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	gsearch "github.com/serpapi/google-search-results-golang"
)

func getAvailabilities() ([]string, error) {
	parameter := map[string]string{
		"engine":        "google_flights",
		"hl":            "fr",
		"gl":            "fr",
		"departure_id":  "NOU",
		"arrival_id":    "TLS",
		"outbound_date": "2024-06-08",
		"currency":      "EUR",
		"type":          "2",
	}

	availabilities := make([]string, 0)

	for _, date := range dates {
		parameter["outbound_date"] = date

		search := gsearch.NewGoogleSearch(parameter, "3cdc9b45fe97c72642a213a24f00c72aa60f0df9c3d7ea6fb224412b558256e3")
		result, err := search.GetJSON()
		if err != nil {
			if strings.Contains(err.Error(), "Google Flights hasn't returned any results for this query") {
				fmt.Printf("No flight for %s\n", date)

				continue
			}

			fmt.Println(err.Error())

			return nil, errors.Wrap(err, "failed to send search")
		}

		json, err := json.Marshal(result)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal to json")
		}

		fmt.Printf("Flights for date %s:\n", date)
		fmt.Println(string(json) + "\n")

		availabilities = append(availabilities, date)
	}

	return availabilities, nil
}
