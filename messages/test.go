package messages

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	AddMessage("test", handleTest)
}

func handleTest(d *discordgo.Session, m *discordgo.MessageCreate) {
	ReplyWith(d, "ayo", m.Reference())
}
