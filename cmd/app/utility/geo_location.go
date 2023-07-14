package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type googleGeoResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

func GetCoordinates(address string) (float64, float64, error) {
	var api = "https://maps.googleapis.com/maps/api/geocode/json"

	response, err := http.Get(fmt.Sprintf("%s?address=%s&key=%s", api, url.QueryEscape(address), "your-api-key"))
	if err != nil {
		return 0, 0, err
	}
	defer response.Body.Close()

	var data googleGeoResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return 0, 0, err
	}

	if len(data.Results) > 0 {
		return data.Results[0].Geometry.Location.Lat, data.Results[0].Geometry.Location.Lng, nil
	}

	return 0, 0, fmt.Errorf("no results")
}
