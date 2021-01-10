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

func updateSession(s *discordgo.Session, m *discordgo.MessageCreate, dice []int, set [][]string) string {
	if _, ok := sessions[m.Author.ID]; ok {
		letter(s, m)
		return ""
	}
	sessions[m.Author.ID] = trick{
		getName(dice, s, m, set),
		m.Author.ID,
		dice,
		"",
		2,
		0,
	}
	archiveJSON(os.Getenv("SK8DICE"), &sessions)
	return getName(dice, s, m, set)
}
