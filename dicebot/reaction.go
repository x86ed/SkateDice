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
	if val, ok := sessions[m.UserID]; ok {
		fmt.Println(m.Emoji.Name)
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
