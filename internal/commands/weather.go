package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func cmdWeather(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Please provide a location. Usage: `!weather [location]`")
		return
	}

	location := strings.Join(args[1:], " ")
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		s.ChannelMessageSend(m.ChannelID, "Weather service is not configured.")
		return
	}

	query := url.QueryEscape(location)
	apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", query, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != 200 {
		s.ChannelMessageSend(m.ChannelID, "Error fetching weather data.")
		return
	}
	defer resp.Body.Close()

	var data struct {
		Name string `json:"name"`
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil || len(data.Weather) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Could not retrieve weather information.")
		return
	}

	message := fmt.Sprintf("Weather in %s: %s, %.1f°C, Humidity: %d%%",
		data.Name, data.Weather[0].Description, data.Main.Temp, data.Main.Humidity)
	s.ChannelMessageSend(m.ChannelID, message)
}
