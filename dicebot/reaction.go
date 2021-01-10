package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Reaction reaction roles
func Reaction(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.UserID == s.State.User.ID {
		return
	}
	if m.Emoji.Name == "👎" {
		fmt.Println("dis")
	} else if m.Emoji.Name == "👍" {
		fmt.Println("like")
	}
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.

	// trigger, render, content, _ := parseTriggers(m.Content)
	// if trigger {
	// 	log.Print(content)
	// 	render(s, m)
	// }
}
