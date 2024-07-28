package messages

import "github.com/bwmarrin/discordgo"

func init() {
	AddMessage("ping", handlePing)
}

func handlePing(d *discordgo.Session, m *discordgo.MessageCreate) {
	d.ChannelMessageSend(m.ChannelID, "pong")
}
