package main

import (
	"bufio"
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

	location := handleInput()
	// GeoCoding: get address long, lat
	geoCodingurl := "https://maps.googleapis.com/maps/api/geocode/json?address=" + location + ",+GB&key=" + apiKey
	geoCode, err := getAPI(geoCodingurl)
	if err != nil {
		fmt.Println("Error:", err)
	}

	lat := geoCode.Results[0].Geometry.Location.Lat
	long := geoCode.Results[0].Geometry.Location.Lng
	fmt.Println("Address:", geoCode.Results[0].FormattedAddress)

	// Weather: get weather data
	weatherAPI := fmt.Sprintf("https://weather.googleapis.com/v1/currentConditions:lookup?key=%v&location.latitude=%v&location.longitude=%v", apiKey, lat, long)
	weather, err := getAPI(weatherAPI)
	if err != nil {
		fmt.Println("Error:", err)
	}
	parseWeather(geoCode.Results[0].FormattedAddress, weather.WeatherResponse)
}

func parseResponse(r io.Reader) (GoogleResponse, error) {
	dec := json.NewDecoder(r)
	var response GoogleResponse
	if err := dec.Decode(&response); err != nil {
		return response, fmt.Errorf("failed to decode response: %w", err)
	}
	return response, nil
}

func parseWeather(address string, weather google.WeatherResponse) {
	fmt.Printf("The weather today in %v is looking %v, with a temperaure of %v and an %v%% chance of rain,", address, weather.WeatherCondition.Type, weather.Temperature.Degrees, weather.Precipitation.Probability.Percent)
}

func getAPI(api string) (GoogleResponse, error) {
	resp, err := http.Get(api)
	if err != nil {
		fmt.Println("Error: weatherAPI")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return parseResponse(resp.Body)
}

func handleInput() string {
	fmt.Println("Enter city name for weather report")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
