package geo

import (
	"encoding/json"
	"net/http"
	"time"
)

type IPInfo struct {
	City    string  `json:"city"`
	Region  string  `json:"region"`
	Country string  `json:"country_name"`
	Lat     float64 `json:"latitude"`
	Lon     float64 `json:"longitude"`
}

func GetIPLocation() (IPInfo, error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://ipapi.co/json")
	if err != nil {
		return IPInfo{}, err
	}
	defer resp.Body.Close()
	var info IPInfo
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&info); err != nil {
		return IPInfo{}, err
	}
	return info, nil
}
