package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	apiKey  string
	apiHost string
)

func init() {
	apiKey = os.Getenv("APIKEY")
	apiHost = os.Getenv("APIHOST")
}

type WeatherRequest struct {
	Cities []string `json:"cities"`
}

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
}

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	http.HandleFunc("/weather", handleWeather)
	log.Fatal(http.ListenAndServe(":8800", nil))
}

func handleWeather(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*") // or specify the allowed origin instead of "*"
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
        return
    }
	
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Invalid request method")
		return
	}

	var request WeatherRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request body")
		return
	}

	weatherData, err := getWeatherData(request.Cities)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "An error occurred")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}

func getWeatherData(cities []string) ([]WeatherData, error) {
	weatherData := make([]WeatherData, len(cities))
	for i, city := range cities {
		temperature, err := getWeather(city)
		if err != nil {
			temperature = 0.0
		}
		weatherData[i] = WeatherData{
			City:        city,
			Temperature: temperature,
		}
	}
	return weatherData, nil
}

func getWeather(city string) (float64, error) {
	url := fmt.Sprintf("https://open-weather13.p.rapidapi.com/city/%s", city)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0.0, err
	}

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("weather request failed with status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0.0, err
	}

	var weatherResponse struct {
		Main struct {
			Temperature float64 `json:"temp"`
		} `json:"main"`
	}

	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return 0.0, err
	}

	return weatherResponse.Main.Temperature, nil
}