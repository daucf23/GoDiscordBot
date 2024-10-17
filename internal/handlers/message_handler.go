package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/daucf23/GoDiscordBot/config"
	"github.com/daucf23/GoDiscordBot/internal/commands"
)

// messageHandler handles incoming messages
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}

		content := strings.TrimPrefix(m.Content, config.BotPrefix)
		content = strings.ToLower(strings.TrimSpace(content))

		commands.CheckUserInput(content, s, m)

	}
}
