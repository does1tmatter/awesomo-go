package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RegisterEvents(d *discordgo.Session) {
	fmt.Println("Registering events")
	d.AddHandler(onClientReady)
	d.AddHandler(onMessageCreate)
}
