package main

import (
	"fmt"

	"github.com/daucf23/GoDiscordBot/config"
	"github.com/daucf23/GoDiscordBot/internal"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	internal.BotStart()
	<-make(chan struct{})
}
