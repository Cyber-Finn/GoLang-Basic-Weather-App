package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const apiKey = "YOUR_API_KEY" //todo: replace this with your own API key from OpenWeatherMap API

// struct based on the response returned from their API
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	var city string
	_, err := fmt.Scanf("Enter the name of the City to check for: %s", &city) // loading up the city - Clang style

	//if there was an error
	if err != nil {
		log.Fatal(err) //this will throw, causing us to skip execution of all the code below(Which is a good thing)
	}

	//build up the URL
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

	//remember that we need the net/http client lib installed to do this: can install by running "go get net/http"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}
	//defer here basically tells the app "remember to close this later - when we're done". Could just close at the end, but deferring close is a nice way to ensure that we dont forget!
	defer resp.Body.Close()

	//var of our struct type
	var weather WeatherResponse

	//if we had any issues, throw an error and exit
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("The current temperature in %s is: %.1f°C\n", city, weather.Main.Temp)

	//just doing this to keep the console window open -> when a console app finishes, it tends to just close the window, without the user seeing the final output
	//this very simple method lets us combat that
	var temp string
	fmt.Scanln(&temp)
}
