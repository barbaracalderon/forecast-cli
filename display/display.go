package display

import (
	"fmt"
	"time"
	"strings"
	"forecast-cli/api"
	"forecast-cli/config"
)

func DisplayWeather(weather *api.WeatherResponse, locationInfo config.LocationInfo) {
	today := weather.Forecast.ForecastDay[0]
	todayDate, _ := time.Parse("2006-01-02", today.Date)

	fmt.Println("┌──────────────────────────────────────────────────────────────────┐")
	fmt.Printf("│ 🌍 Location: %s, %s, %s\n", locationInfo.City, locationInfo.Region, locationInfo.Country)
	fmt.Printf("│ 📅 Date: %s | %s\n", todayDate.Format("02/01"), todayDate.Format("Monday"))
	fmt.Println("├──────────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ 🌡️ Current: %s, %.0f°C (Feels like %.0f°C)\n", strings.ToLower(weather.Current.Condition.Text), 
    today.Day.AvgTempC, weather.Current.FeelsLikeC)
	fmt.Printf("│ 🔽 Min: %.0f°C | 🔼 Max: %.0f°C\n", today.Day.MinTempC, today.Day.MaxTempC)
	fmt.Printf("│ 💧 Humidity: %d%%\n", weather.Current.Humidity)
	fmt.Println("└──────────────────────────────────────────────────────────────────┘")
	fmt.Println()

	fmt.Println("📅 6-Day Forecast:")
	fmt.Println("┌───────────┬───────────┬─────────────┬──────────────────────┬──────────┬──────────┐")
	fmt.Println("│   Date    │ Temp (°C) │ Rain Chance │      Condition       │ Sunrise  │  Sunset  │")
	fmt.Println("├───────────┼───────────┼─────────────┼──────────────────────┼──────────┼──────────┤")

	for i := range weather.Forecast.ForecastDay {
		day := weather.Forecast.ForecastDay[i]
		date, _ := time.Parse("2006-01-02", day.Date)
		condition := strings.ToLower(day.Day.Condition.Text)

		if len(condition) > 20 {
			condition = condition[:20] + "..."
		}

		dateStr := fmt.Sprintf("%s %s", date.Format("02/01"), date.Format("Mon"))
		rainChance := fmt.Sprintf("%.0f%%", day.Day.TotalPrecipMM)

		fmt.Printf(
			"│ %-9s │ %-9.0f │ %-11s │ %-20s │ %-8s │ %-8s │\n",
			dateStr,
			day.Day.AvgTempC,
			rainChance,
			condition,
			day.Astro.Sunrise,
			day.Astro.Sunset,
		)
	}

	fmt.Println("└───────────┴───────────┴─────────────┴──────────────────────┴──────────┴──────────┘")
}