package commands

import "github.com/bwmarrin/discordgo"

func cmdPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "pong")
}
