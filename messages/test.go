package messages

import "github.com/bwmarrin/discordgo"

func init() {
	AddMessage("test", handleTest)
}

func handleTest(d *discordgo.Session, m *discordgo.MessageCreate) {
	d.ChannelMessageSend(m.ChannelID, "ayo")
}
