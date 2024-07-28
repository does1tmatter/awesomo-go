package events

import (
	"awesomo-go/messages"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	addHandler(onMessageCreate)
}

func onMessageCreate(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	go handleTestMessage(discord, message)
}

func handleTestMessage(d *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("Message received: %v\n", m.Content)

	callback := messages.GetMessage(m.Content)
	if callback != nil {
		callback(d, m)
	}
}
