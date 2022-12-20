package main

import (
	"fmt"

	"github.com/shashank/discord-pingbot/bot"
	"github.com/shashank/discord-pingbot/config"
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
