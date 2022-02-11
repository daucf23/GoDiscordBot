package main

import (
	"fmt"

	"github.com/daucf23/GoDiscordBot/bot"
	"github.com/daucf23/GoDiscordBot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
