package messages

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var allMessages map[string]func(
	*discordgo.Session,
	*discordgo.MessageCreate,
) = make(map[string]func(
	*discordgo.Session,
	*discordgo.MessageCreate,
))

func AddMessage(key string, callback func(*discordgo.Session, *discordgo.MessageCreate)) {
	allMessages[key] = callback
}

func GetMessage(key string) func(*discordgo.Session, *discordgo.MessageCreate) {
	callback, exists := allMessages[key]
	if !exists {
		return nil
	}

	return callback
}

func ReplyWith(d *discordgo.Session, content string, m *discordgo.MessageReference) {
	message, err := d.ChannelMessageSendReply(m.ChannelID, content, m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Responded: %v\n\n", message.Content)
}
