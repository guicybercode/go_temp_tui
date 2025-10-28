package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Current struct {
	Name string `json:"city_name"`
	Data []struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"app_temp"`
		Weather   struct {
			Main        string `json:"weather"`
			Icon        string `json:"icon"`
			Description string `json:"description"`
		} `json:"weather"`
		Wind struct {
			Speed float64 `json:"wind_spd"`
		} `json:"wind"`
		Dt int64 `json:"ts"`
	} `json:"data"`
}

type Forecast struct {
	Data []struct {
		Dt      int64   `json:"ts"`
		Temp    float64 `json:"temp"`
		Weather struct {
			Main string `json:"weather"`
			Icon string `json:"icon"`
		} `json:"weather"`
	} `json:"data"`
}

type Client struct {
	APIKey string
	Units  string
	HTTP   http.Client
}

func NewClient(key, units string) Client {
	return Client{APIKey: key, Units: units, HTTP: http.Client{Timeout: 7 * time.Second}}
}

func (c Client) CurrentByCoords(lat, lon float64, lang string) (Current, error) {
	endpoint := "https://api.weatherbit.io/v2.0/current"
	q := url.Values{}
	q.Set("lat", fmt.Sprintf("%f", lat))
	q.Set("lon", fmt.Sprintf("%f", lon))
	q.Set("key", c.APIKey)
	q.Set("units", c.Units)
	if lang != "" {
		q.Set("lang", lang)
	}
	resp, err := c.HTTP.Get(endpoint + "?" + q.Encode())
	if err != nil {
		return Current{}, err
	}
	defer resp.Body.Close()
	var cur Current
	if err := json.NewDecoder(resp.Body).Decode(&cur); err != nil {
		return Current{}, err
	}
	return cur, nil
}

func (c Client) ForecastByCoords(lat, lon float64, lang string) (Forecast, error) {
	endpoint := "https://api.weatherbit.io/v2.0/forecast/hourly"
	q := url.Values{}
	q.Set("lat", fmt.Sprintf("%f", lat))
	q.Set("lon", fmt.Sprintf("%f", lon))
	q.Set("key", c.APIKey)
	q.Set("units", c.Units)
	q.Set("hours", "12")
	if lang != "" {
		q.Set("lang", lang)
	}
	resp, err := c.HTTP.Get(endpoint + "?" + q.Encode())
	if err != nil {
		return Forecast{}, err
	}
	defer resp.Body.Close()
	var fc Forecast
	if err := json.NewDecoder(resp.Body).Decode(&fc); err != nil {
		return Forecast{}, err
	}
	return fc, nil
}
