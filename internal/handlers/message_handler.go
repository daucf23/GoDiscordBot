package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/daucf23/GoDiscordBot/config"
	"github.com/daucf23/GoDiscordBot/internal/commands"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}
	if m.Author.ID == BotID {
		return
	}

	content := strings.TrimPrefix(m.Content, config.BotPrefix)
	content = strings.TrimSpace(content)

	// Split content into command and arguments
	args := strings.Fields(content)
	if len(args) == 0 {
		return
	}

	command := strings.ToLower(args[0])
	commands.CheckUserInput(command, args, s, m)
}
