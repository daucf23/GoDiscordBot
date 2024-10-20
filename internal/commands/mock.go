package commands

import (
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
)

func cmdMock(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Please provide text to mock. Usage: `!mock [text]`")
		return
	}

	input := strings.Join(args[1:], " ")
	mocked := mockText(input)
	s.ChannelMessageSend(m.ChannelID, mocked)
}

func mockText(s string) string {
	var result strings.Builder
	upper := false
	for _, c := range s {
		ch := string(c)
		if upper {
			result.WriteString(strings.ToUpper(ch))
		} else {
			result.WriteString(strings.ToLower(ch))
		}
		if unicode.IsLetter(c) {
			upper = !upper
		}
	}
	return result.String()
}
