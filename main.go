package main

import (
	"flag"
	"fmt"
	"os"

	"forecast-cli/api"
	"forecast-cli/config"
	"forecast-cli/display"
)

func main() {
	locationFlag := flag.String("location", "", "Specify a custom location for the weather forecast")
	flag.Parse()

	if len(os.Args) > 0 && (os.Args[0] == "forecast" || (len(os.Args) > 1 && os.Args[1] == "forecast")) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("error loading config:", err)
			os.Exit(1)
		}

		var location string
		if *locationFlag != "" {
			location = *locationFlag
		} else {
			location = cfg.IPInfo.City
		}

		weather, err := api.FetchWeather(cfg.APIKey, location)
		if err != nil {
			fmt.Println("error fetching weather:", err)
			os.Exit(1)
		}

		displayLocation := config.LocationInfo{
			City:    weather.Location.Name,
			Region:  weather.Location.Region,
			Country: weather.Location.Country,
		}

		display.DisplayWeather(weather, displayLocation)
	} else {
		fmt.Println("Usage: forecast [--location <city>]")
		flag.PrintDefaults()
		os.Exit(1)
	}
}