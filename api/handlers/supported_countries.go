package handlers

import (
	"encoding/json"
	"net/http"
)

type Countries struct {
	AvailableCountries []Country `json:"available_countries,omitempty"`
}

type Country struct {
	Name string      `json:"name,omitempty"`
	Code interface{} `json:"code,omitempty"`
}

func SupportedCountryHandler(w http.ResponseWriter, r *http.Request) {
	countries := map[string]interface{}{
		"Argentina":      "AR",
		"Australia":      "AU",
		"Austria":        "AT",
		"Belgium":        "BE",
		"Brazil":         "BR",
		"Canada":         "CA",
		"Chile":          "CL",
		"Colombia":       "CO",
		"Czechia":        "CZ",
		"Denmark":        "DK",
		"Egypt":          "EG",
		"Finland":        "FI",
		"France":         "FR",
		"Germany":        "DE",
		"Greece":         "GR",
		"Hong Kong":      "HK",
		"Hungary":        "HU",
		"India":          "IN",
		"Indonesia":      "ID",
		"Ireland":        "IE",
		"Israel":         "IL",
		"Italy":          "IT",
		"Japan":          "JP",
		"Kenya":          "KE",
		"Malaysia":       "MY",
		"Mexico":         "MX",
		"Netherlands":    "NL",
		"New Zealand":    "NZ",
		"Nigeria":        "NG",
		"Norway":         "NO",
		"Peru":           "PE",
		"Philippines":    "PH",
		"Poland":         "PL",
		"Portugal":       "PT",
		"Romania":        "RO",
		"Russia":         "RU",
		"Saudi Arabia":   "SA",
		"Singapore":      "SG",
		"South Africa":   "ZA",
		"South Korea":    "KR",
		"Spain":          "ES",
		"Sweden":         "SE",
		"Switzerland":    "CH",
		"Taiwan":         "TW",
		"Thailand":       "TH",
		"Turkey":         "TR",
		"Ukraine":        "UA",
		"United Kingdom": "GB",
		"United States":  "US",
		"Vietnam":        "VN",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	countryR := make([]Country, 0, len(countries))

	for i, j := range countries {
		countryR = append(countryR, Country{Name: i, Code: j})
	}

	response := Countries{
		AvailableCountries: countryR,
	}

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

}
