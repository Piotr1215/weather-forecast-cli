// Package main provides a command-line application that retrieves and displays
// weather information including current conditions, a three-day forecast, and
// sunrise/sunset times for a specified location. The data is fetched from the
// weatherapi.com service using an API key.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Conditions string `json:"text"`
		} `json:"condition"`
	}
	Forecast struct {
		ForecastDay []struct {
			Date string `json:"date"`
			Hour []struct {
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Conditions string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
			Cycle struct {
				Sunrise string `json:"sunrise"`
				Sunset  string `json:"sunset"`
			} `json:"astro"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	k := os.Getenv("WEATHER_API_KEY")
	if k == "" {
		fmt.Println("WEATHER_API_KEY is not set")
		os.Exit(1)
	}
	l := os.Getenv("WEATHER_LOCATION")

	if l == "" {
		fmt.Println("WEATHER_LOCATION is not set")
		os.Exit(1)
	}

	// pull data from the weather api
	// free tier only allows 3 days of forecast and 1 million requests per month
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=" + k + "&q=" + l + "&days=3&aqi=no&alerts=no")
	if err != nil {
		log.Fatalf("Error connecting to the weather API:\n%v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Expected 200 but got: ", res.Status+"\n")
		fmt.Println("Response Headers:")
		for key, values := range res.Header {
			// Join the slice of values for each header key into a single string
			valueString := strings.Join(values, ", ")
			fmt.Printf("%s: %s\n", key, valueString)
		}
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body:\n%v", err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON:\n%v", err)
	}
	fmt.Println("\nWeather in", weather.Location.Name, "now is:", weather.Current.TempC, "degrees Celsius and", weather.Current.Condition.Conditions)

	fmt.Println("\nForecast for next days at 11 AM:")
	layout := "2006-01-02 15:04" // Go reference time format
	for _, day := range weather.Forecast.ForecastDay {
		for _, hour := range day.Hour {
			parsedTime, errr := time.Parse(layout, hour.Time)
			if errr != nil {
				fmt.Println("Error parsing time:", err)
				continue
			}
			if parsedTime.Hour() == 11 && parsedTime.Day() != time.Now().Day() {
				fmt.Println("- Temperature on", day.Date, "at 11 AM will be", hour.TempC, "°C", "with conditions", hour.Condition.Conditions)
			}
		}
	}

	fmt.Println("\nSunrise and sunset and day length:")
	fmt.Println("- Today sun rose at", weather.Forecast.ForecastDay[0].Cycle.Sunrise, "and will set at", weather.Forecast.ForecastDay[0].Cycle.Sunset)

	sunrise := weather.Forecast.ForecastDay[0].Cycle.Sunrise
	sunset := weather.Forecast.ForecastDay[0].Cycle.Sunset

	hours, minutes, err := calculateDaylightDuration(sunrise, sunset)
	if err != nil {
		fmt.Println("Error calculating daylight duration:", err)
	} else {

		fmt.Printf("- Time between sunrise and sunset is: %d hours and %d minutes\n", hours, minutes)
	}

}

// calculateDaylightDuration calculates the duration of daylight based on sunrise and sunset times.
// Returns the duration in hours and minutes.
func calculateDaylightDuration(sunrise, sunset string) (int, int, error) {
	const timeLayout = "03:04 PM"
	sunriseTime, err := time.Parse(timeLayout, sunrise)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting sunrise string to time: %v", err)
	}

	sunsetTime, err := time.Parse(timeLayout, sunset)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting sunset string to time: %v", err)
	}

	duration := sunsetTime.Sub(sunriseTime)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - (hours * 60)
	return hours, minutes, nil
}
