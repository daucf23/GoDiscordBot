package commands

import (
	"github.com/bwmarrin/discordgo"
)

func CheckUserInput(command string, args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	switch command {
	case "ping":
		cmdPing(s, m)
	case "help":
		cmdHelp(s, m)
	case "weather":
		cmdWeather(s, m, args)
	case "joke":
		cmdJoke(s, m)
	case "fact":
		cmdFact(s, m)
	case "8ball":
		cmd8Ball(s, m, args)
	case "reverse":
		cmdReverse(s, m, args)
	case "mock":
		cmdMock(s, m, args)
	case "ask":
		cmdAsk(s, m, args)
	case "quote":
		cmdQuote(s, m)
	// Add more cases for your other commands
	default:
		s.ChannelMessageSend(m.ChannelID, "Unknown command. Type `!help` for a list of commands.")
	}
}
