package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var handlers []interface{} = make([]interface{}, 0, 10)

func addHandler(handler interface{}) {
	handlers = append(handlers, handler)
}

func RegisterEvents(d *discordgo.Session) {
	fmt.Println("Registering events")
	for _, v := range handlers {
		d.AddHandler(v)
	}
}
