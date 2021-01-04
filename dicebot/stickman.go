package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var prefixes = []string{
	"!sd",
	"!dice",
	"!skatedice",
	"!skatefairy",
	"!sf",
}

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	out := "```yaml\n" +
		`Triggers:\n* 
` + strings.Join(prefixes, "\n* ") +
		`\nCommands:
* !sd help/list - show this 
* !sd roll/regular/reg - roll the dice for normal tricks
* !sd bowl/vert/pool/trans (coming soon) - roll the dice for transition tricks
* !sd park/ledge/curb (coming soon) - roll the dice for grinds
* !sd letter - take a letter and forfeit the round
* !sd reset - remove all of your levels
` + "```"
	s.ChannelMessageSend(m.ChannelID, out)
}

func rollDice(set [][]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i0 := r1.Intn(6)
	i1 := r1.Intn(6)
	i2 := r1.Intn(6)
	i3 := r1.Intn(6)
	i4 := r1.Intn(6)
	i5 := r1.Intn(6)
	lvl := getDif(m)
	if i1 == 0 && i2 == 0 && i3 == 0 && i0 == 0 {
		img := lucky[i4+i5%len(lucky)]
		f := []*discordgo.MessageEmbedField{
			{
				Name:  "difficulty",
				Value: lvl.name,
			},
		}
		e := discordgo.MessageEmbed{
			Color:       lvl.color,
			Title:       "Randomized",
			Description: "pick a random trick or roll again",
			URL:         "https://discord.gg/zq3fyV",
			Author: &discordgo.MessageEmbedAuthor{
				URL:     "https://discord.gg/zq3fyV",
				Name:    "GrossoBot",
				IconURL: "https://i.ibb.co/4RBtbVC/grossobot.gif",
			},
			Fields: f,
		}
		e.Image = &discordgo.MessageEmbedImage{
			URL: img,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &e)
		return
	}
	if i1 == 5 && i2 == 5 && i3 == 5 && i0 == 5 {
		desc, img := giveLetter(s, m, sucky[i4+i5%len(sucky)])
		f := []*discordgo.MessageEmbedField{
			{
				Name:  "difficulty",
				Value: lvl.name,
			},
		}
		e := discordgo.MessageEmbed{
			Color:       lvl.color,
			Title:       "Take a Letter",
			Description: desc,
			URL:         "https://discord.gg/zq3fyV",
			Author: &discordgo.MessageEmbedAuthor{
				URL:     "https://discord.gg/zq3fyV",
				Name:    "GrossoBot",
				IconURL: "https://i.ibb.co/4RBtbVC/grossobot.gif",
			},
			Fields: f,
		}
		if len(img) > 1 {
			e.Image = &discordgo.MessageEmbedImage{
				URL: img,
			}
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &e)
		return
	}
	img := lucky[i4+i5%len(lucky)]
	desc := ""
	if len(set) < 1 {
		set = qs[lvl.name]
	}
	if i0 != 0 && i0 != 5 {
		desc = set[0][i0] + " "

	}
	if i1 != 0 && i1 != 5 {
		desc += set[1][i1] + " "

	}
	if i2 != 0 && i2 != 5 {
		desc += set[2][i2] + " "

	}
	if i3 != 0 && i3 != 5 {
		desc += set[3][i3] + " "

	}
	if lvl.color == 0xFF00FF {
		if i4 != 0 && i4 != 5 {
			desc += set[4][i4] + " "

		}
		if i5 != 0 && i5 != 5 {
			desc += set[5][i5] + " "

		}
	}
	f := []*discordgo.MessageEmbedField{
		{
			Name:  "difficulty",
			Value: lvl.name,
		},
	}
	log.Print(desc)
	e := discordgo.MessageEmbed{
		Color:       lvl.color,
		Title:       fmt.Sprintf("Do a %s.", getTitle(desc)),
		Description: desc,
		URL:         "https://discord.gg/zq3fyV",
		Author: &discordgo.MessageEmbedAuthor{
			URL:     "https://discord.gg/zq3fyV",
			Name:    "GrossoBot",
			IconURL: "https://i.ibb.co/4RBtbVC/grossobot.gif",
		},
		Fields: f,
	}
	if len(img) > 1 {
		e.Image = &discordgo.MessageEmbedImage{
			URL: img,
		}
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &e)
}

func regDice(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print("regdice")
	rollDice([][]string{}, s, m)
}

func reset(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := removeAllLetters(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error resseting the user!")
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@!%s>'s game was reset!", m.Author.ID))
}

func letter(s *discordgo.Session, m *discordgo.MessageCreate) {
	desc, img := giveLetter(s, m, "https://i.ibb.co/WPqS8Jw/jarne.gif")
	e := discordgo.MessageEmbed{
		Color:       0xFFFF00,
		Title:       "Take the L",
		Description: desc,
		URL:         "https://discord.gg/zq3fyV",
		Author: &discordgo.MessageEmbedAuthor{
			URL:     "https://discord.gg/zq3fyV",
			Name:    "GrossoBot",
			IconURL: "https://i.ibb.co/4RBtbVC/grossobot.gif",
		},
	}
	e.Image = &discordgo.MessageEmbedImage{
		URL: img,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &e)
}

func getHandler(s string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	handlers := map[string]func(s *discordgo.Session, m *discordgo.MessageCreate){
		"reg":  regDice,
		"roll": regDice,
		// "vert": vertDice,
		// "bowl": vertDice,
		// "pool": vertDice,
		// "trans": vertDice,
		// "park": ledDice,
		// "ledge": ledDice,
		// "curb": ledDice,
		"reset":  reset,
		"letter": letter,
		"help":   help,
		"list":   help,
	}
	for i, v := range handlers {
		if strings.HasPrefix(s, i) {
			return v
		}
	}
	return help
}

func parseTriggers(s string) (bool, func(s *discordgo.Session, m *discordgo.MessageCreate), string, error) {
	for _, v := range prefixes {
		if strings.HasPrefix(s, v) {
			n := strings.Replace(s, v+" ", "", -1)
			h := getHandler(n)
			return true, h, s, nil
		}
	}
	return false, func(s *discordgo.Session, m *discordgo.MessageCreate) {}, "", errors.New("dice")
}

// Stickman rolls the dice for you.
func Stickman(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	trigger, render, content, _ := parseTriggers(m.Content)
	if trigger {
		log.Print(content)
		render(s, m)
	}
}
