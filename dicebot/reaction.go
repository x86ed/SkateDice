package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Reaction reaction roles
func Reaction(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	log.Println("userID ", m.UserID, m.Emoji.User.ID, m.Emoji.User.Username)
	log.Println("State UserID", s.State.User.ID, s.State.User.Username)
	if m.UserID == s.State.User.ID {
		return
	}
	if val, ok := sessions[m.UserID]; ok {
		log.Printf(m.Emoji.Name)
		if m.MessageID == val.Msg {
			if m.Emoji.Name == "😡" {
				fmt.Println("dis")
				val.Score--
				if val.Score <= 0-val.Needed {
					giveLetterR(s, m, sucky[val.Dice[4]+val.Dice[5]%len(sucky)])
				}
			} else if m.Emoji.Name == "😍" {
				fmt.Println("like")
				val.Score++
				if val.Score >= val.Needed {
					addDif(s, m)
				}
			}
		}
	}
}
