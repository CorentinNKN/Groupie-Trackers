package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Point GPS
type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Transforme une ville en coordonnées GPS
func GeocodeCity(city string) (*GeoPoint, error) {
	apiURL := fmt.Sprintf(
		"https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1",
		url.QueryEscape(city),
	)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Obligatoire pour Nominatim
	req.Header.Set("User-Agent", "concert-map-app")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("ville introuvable")
	}

	lat, _ := strconv.ParseFloat(data[0].Lat, 64)
	lng, _ := strconv.ParseFloat(data[0].Lon, 64)

	return &GeoPoint{
		Lat: lat,
		Lng: lng,
	}, nil
}
