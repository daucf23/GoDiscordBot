package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/rand"
)

var responses = []string{
	"It is certain.",
	"Reply hazy, try again.",
	"Don't count on it.",
	"Yes, definitely.",
	"My reply is no.",
	// Add more responses
}

func cmd8Ball(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You need to ask a question. Usage: `!8ball [question]`")
		return
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	response := responses[rand.Intn(len(responses))]
	s.ChannelMessageSend(m.ChannelID, response)
}
