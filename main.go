package main

import (
	"fmt"

	"awesomo-go/env"
	"awesomo-go/events"
	keepalive "awesomo-go/keep-alive"

	"github.com/bwmarrin/discordgo"
)

func main() {
	env.InitEnv()
	fmt.Println("Initiating client")

	discord, err := discordgo.New("Bot " + env.GetEnv("AUTH_KEY"))
	if err != nil {
		keepalive.Start(err)
	}

	events.RegisterEvents(discord)

	fmt.Println("Starting discord session")
	err = discord.Open()
	if err != nil {
		keepalive.Start(err)
	}

	defer disconnect(discord)

	keepalive.Start()
}

func disconnect(d *discordgo.Session) {
	fmt.Println("Disconnecting session")

	err := d.Close()
	if err != nil {
		fmt.Printf("Error while disconnecting: %s", err)
		return
	}

	fmt.Println("Session disconnected succesfully")
}
