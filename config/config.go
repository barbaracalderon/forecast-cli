package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
	IPInfo LocationInfo
}

type LocationInfo struct {
    City    string `json:"city"`
    Region  string `json:"region"`
    Country string `json:"country"`
}

func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %v", err)
	}

	envPath := filepath.Join(homeDir, ".config", "forecast-cli", ".env")

	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	apiKey := os.Getenv("WEATHERAPI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing required environment variable: WEATHERAPI_API_KEY")
	}

	ipInfo, err := getLocationByIP()
	if err != nil {
		return nil, fmt.Errorf("error fetching location: %v", err)
	}

	return &Config{
		APIKey: apiKey,
		IPInfo: ipInfo,
	}, nil
}

func getLocationByIP() (LocationInfo, error) {
	resp, err := http.Get("https://ipinfo.io")
	if err != nil {
		return LocationInfo{}, fmt.Errorf("error fetching IP info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationInfo{}, fmt.Errorf("IP info request failed with status: %s", resp.Status)
	}

	var locationInfoFromIP LocationInfo
	if err := json.NewDecoder(resp.Body).Decode(&locationInfoFromIP); err != nil {
		return LocationInfo{}, fmt.Errorf("error decoding IP info response: %v", err)
	}

	if locationInfoFromIP.City == "" || locationInfoFromIP.Region == "" || locationInfoFromIP.Country == "" {
		return LocationInfo{}, fmt.Errorf("could not determine location from IP address")
	}

	return locationInfoFromIP, nil
}