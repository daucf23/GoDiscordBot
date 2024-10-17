package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var BotID string

func GetBotID(u *discordgo.User, goBot *discordgo.Session) string {
	BotID = u.ID
	return BotID
}
