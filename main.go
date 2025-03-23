package main

import (
	"fmt"
	"os"

	"forecast-cli/api"
	"forecast-cli/config"
	"forecast-cli/display"
)

func main() {
	if len(os.Args) > 0 && (os.Args[0] == "forecast" || (len(os.Args) > 1 && os.Args[1] == "forecast")) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("error loading config:", err)
			os.Exit(1)
		}

		weather, err := api.FetchWeather(cfg.APIKey, cfg.IPInfo.City)
		if err != nil {
			fmt.Println("error fetching weather:", err)
			os.Exit(1)
		}

		display.DisplayWeather(weather, cfg.IPInfo)
	} else {
		fmt.Println("Usage: forecast")
		os.Exit(1)
	}
}