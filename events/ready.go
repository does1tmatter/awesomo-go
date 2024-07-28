package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	addHandler(onClientReady)
}

func onClientReady(_ *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("Discord client ready.\nConnected as %v\n\n", r.User.String())
}
