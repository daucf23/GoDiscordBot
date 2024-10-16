package main

import (
	"fmt"

	"github.com/daucf23/GoDiscordBot/internal"
)

func main() {
	err := internal.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	internal.BotStart()
	<-make(chan struct{})
}
