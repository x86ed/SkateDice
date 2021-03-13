package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type trick struct {
	Name   string
	User   string
	Dice   []int
	Msg    string
	Needed int
	Score  int
}

var sessions = map[string]trick{}

func updateSession(s *discordgo.Session, m *discordgo.MessageCreate, Need int, Score int, dice []int, set [][]string) string {
	if _, ok := sessions[m.Author.ID]; ok {
		letter(s, m)
		return ""
	}
	sessions[m.Author.ID] = trick{
		getName(dice, s, m, set),
		m.Author.ID,
		dice,
		"",
		Need,
		Score,
	}
	archiveJSON(os.Getenv("SK8DICE"), &sessions)
	return getName(dice, s, m, set)
}

func updateSessionR(s *discordgo.Session, m *discordgo.MessageReactionAdd, Need int, Score int, dice []int, set [][]string) string {
	om, _ := s.ChannelMessage(m.ChannelID, m.MessageID)
	mc := &discordgo.MessageCreate{Message: om}
	if _, ok := sessions[om.Author.ID]; ok {
		letter(s, mc)
		return ""
	}
	sessions[om.Author.ID] = trick{
		getName(dice, s, mc, set),
		om.Author.ID,
		dice,
		"",
		Need,
		Score,
	}
	archiveJSON(os.Getenv("SK8DICE"), &sessions)
	return getName(dice, s, mc, set)
}
