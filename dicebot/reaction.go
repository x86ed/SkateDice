package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Reaction reaction roles
func Reaction(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	log.Printf("state %+v  message %+v\n", s, m)
	mem, _ := s.GuildMember(m.GuildID, m.UserID)
	mem2, _ := s.GuildMember(m.GuildID, s.State.User.ID)
	log.Printf("Members %s %s\n ", mem.User.Username, mem2.User.Username)
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
