package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/randleem/go_weather_cli/google"
)

type GoogleResponse struct {
	google.WeatherResponse
	google.GeocodeResponse
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("GOOGLE_GEO_CODING_API_KEY")

	// GeoCoding: get address long, lat
	geoCodingurl := "https://maps.googleapis.com/maps/api/geocode/json?address=131+pineapple+road,+stirchley,+birmingham,+GB&key=" + apiKey
	geoCodeRes, err := http.Get(geoCodingurl)
	if err != nil {
		fmt.Println("Error: weatherAPI")
	}
	defer geoCodeRes.Body.Close()
	if geoCodeRes.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", geoCodeRes.StatusCode, geoCodeRes.Status)
	}
	fmt.Println(geoCodeRes.Status)

	geoCode, err := parseResponse(geoCodeRes.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	lat := geoCode.Results[0].Geometry.Location.Lat
	long := geoCode.Results[0].Geometry.Location.Lng
	fmt.Println("Address:", geoCode.Results[0].FormattedAddress)

	// Weather: get weather data
	resp, err := http.Get(fmt.Sprintf("https://weather.googleapis.com/v1/currentConditions:lookup?key=%v&location.latitude=%v&location.longitude=%v", apiKey, lat, long))
	if err != nil {
		fmt.Println("Error: weatherAPI")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	weather, err := parseResponse(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("The weather today at %v is looking %v, with a temperaure of %v and a %v%% chance of rain,", geoCode.Results[0].FormattedAddress, weather.WeatherCondition.Type, weather.Temperature.Degrees, weather.Precipitation.Probability.Percent)
}

func parseResponse[T GoogleResponse](r io.Reader) (T, error) {
	dec := json.NewDecoder(r)
	var response T
	if err := dec.Decode(&response); err != nil {
		return response, fmt.Errorf("failed to decode response: %w", err)
	}
	return response, nil
}

// 131+pineapple+road,+stirchley,+birmingham,+GB
// func parseAddress(string) string {
// }
