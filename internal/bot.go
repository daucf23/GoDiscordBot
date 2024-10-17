package internal

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/daucf23/GoDiscordBot/config"
	"github.com/daucf23/GoDiscordBot/internal/handlers"
)

var BotID string

// BotStart initializes and starts the Discord bot
func BotStart() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = handlers.GetBotID(u, goBot)

	goBot.AddHandler(handlers.MessageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running or whatever.")
}
