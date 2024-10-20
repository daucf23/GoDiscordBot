package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func cmdReverse(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Please provide text to reverse. Usage: `!reverse [text]`")
		return
	}

	input := strings.Join(args[1:], " ")
	reversed := reverseString(input)
	s.ChannelMessageSend(m.ChannelID, reversed)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
