package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const weatherAPI = "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m,wind_speed_10m"

type CurrentWeather struct {
	Temperature float32 `json:"temperature_2m"`
	WindSpeed   float32 `json:"wind_speed_10m"`
}
type Weather struct {
	Latitude  float32
	Longitude float32
	Current   CurrentWeather
}

func main() {
	resp, err := http.Get(weatherAPI)
	if err != nil {
		fmt.Println("Error: weatherAPI")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	dec := json.NewDecoder(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var weather Weather
	if err := dec.Decode(&weather); err != nil {
		fmt.Print("Error: ", err)
	}
	fmt.Println(weather.Current.Temperature, weather.Current.WindSpeed)
}

// Get weather from api
