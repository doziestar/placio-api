package utility

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"io"
	"log"
	"net/http"
	"net/url"
)

const MapboxAPIKey = "pk.eyJ1IjoiZG96aWVzdGFyIiwiYSI6ImNsazJnNnUzbjBlN20zZXAycXo1NXlka3oifQ.e4rvqdG-6RHXNCbX-s1e3g"

type MapboxGeoResponse struct {
	Type     string `json:"type"`
	Features []struct {
		Type     string `json:"type"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

func GetCoordinates(address string) (MapboxGeoResponse, error) {
	var api = "https://api.mapbox.com/geocoding/v5/mapbox.places/"

	response, err := http.Get(fmt.Sprintf("%s%s.json?access_token=%s", api, url.QueryEscape(address), MapboxAPIKey))
	if err != nil {
		sentry.CaptureException(err)
		log.Println("error getting coordinates", err.Error())
		return MapboxGeoResponse{}, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}

	//bodyString := string(bodyBytes)
	//log.Println("API Response", bodyString)

	var data MapboxGeoResponse
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		sentry.CaptureException(err)
		log.Println("error decoding response", err.Error())
		return MapboxGeoResponse{}, err
	}

	if len(data.Features) > 0 {
		return data, nil
	}

	return MapboxGeoResponse{}, err
}
