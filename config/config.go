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
	IPInfo IPInfoResponse
}

type IPInfoResponse struct {
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

func getLocationByIP() (IPInfoResponse, error) {
	resp, err := http.Get("https://ipinfo.io")
	if err != nil {
		return IPInfoResponse{}, fmt.Errorf("error fetching IP info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return IPInfoResponse{}, fmt.Errorf("IP info request failed with status: %s", resp.Status)
	}

	var ipInfo IPInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		return IPInfoResponse{}, fmt.Errorf("error decoding IP info response: %v", err)
	}

	if ipInfo.City == "" || ipInfo.Region == "" || ipInfo.Country == "" {
		return IPInfoResponse{}, fmt.Errorf("could not determine location from IP address")
	}

	return ipInfo, nil
}