package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
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
` + strings.Join(prefixes, "\n* ") + "\n" +
		`Commands:
* !sd help/list - show this 
* !sd roll/regular/reg - roll the dice for normal tricks
* !sd bowl/vert/pool/trans - roll the dice for transition tricks
* !sd park/ledge/curb - roll the dice for grinds
* !sd jb/malto/manual - roll the dice for malto manual shit
* !sd letter - take a letter and forfeit the round
* !sd reset - remove all of your levels
* !sd submit + a video - a submission to judge
` + "```"
	s.ChannelMessageSend(m.ChannelID, out)
}

func sub(s *discordgo.Session, m *discordgo.MessageCreate) {
	if val, ok := sessions[m.Author.ID]; ok {
		val.Msg = m.Message.ID
		sessions[m.Author.ID] = val
		archiveJSON(os.Getenv("SK8DICE"), &sessions)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@!%s> Just submitted a %s. Vote by reacting to it with 😍 or 😡", m.Author.ID, val.Name))
	}
}

func getName(dice []int, s *discordgo.Session, m *discordgo.MessageCreate, set [][]string) string {
	if len(dice) < 5 {
		return ""
	}
	lvl := getDif(m)
	if len(set) < 1 {
		set = qs[lvl.name]
	}
	desc := ""
	if len(set) < 1 {
		set = qs[lvl.name]
	}
	if dice[0] != 0 && dice[0] != 5 {
		desc = set[0][dice[0]] + " "

	}
	if dice[1] != 0 && dice[1] != 5 {
		desc += set[1][dice[1]] + " "

	}
	if dice[2] != 0 && dice[2] != 5 {
		desc += set[2][dice[2]] + " "

	}
	if dice[3] != 0 && dice[3] != 5 {
		desc += set[3][dice[3]] + " "

	}
	if lvl.color == 0xFF00FF {
		if dice[4] != 0 && dice[4] != 5 {
			desc += set[4][dice[4]] + " "

		}
		if dice[5] != 0 && dice[5] != 5 {
			desc += set[5][dice[5]] + " "

		}
	}
	return desc
}

func getImg(m *discordgo.MessageCreate, i0, i1, i2, i3, i4, i5 int, set [][]string) string {
	lvl := getDif(m)
	if len(set) < 1 {
		set = qs[lvl.name]
	}
	var v0, v1, v2, v3, v4, v5, output string
	if i0 > 0 && i0 < 5 {
		v0 = set[0][i0]
	}
	if i1 > 0 && i1 < 5 {
		v1 = set[1][i1]
	}
	if i2 > 0 && i2 < 5 {
		v2 = set[2][i2]
	}
	if i3 > 0 && i3 < 5 {
		v3 = set[3][i3]
	}
	if i4 > 0 && i4 < 5 {
		v4 = set[4][i4]
	}
	if i5 > 0 && i5 < 5 {
		v5 = set[5][i5]
	}

	if v0 == "fakie" {
		output = "https://i.ibb.co/d4NfG6z/fakie.gif"
	}

	if v0 == "nollie" {
		output = "https://i.ibb.co/PYLqb2X/nollie.gif"
	}

	if v0 == "switch" {
		output = "https://i.ibb.co/VNHBdH5/switch.gif"
	}

	if v2 == "shuv" {
		output = ""
	}

	if v3 == "kickflip" {
		output = "https://i.ibb.co/r6RNQP4/flip.gif"
		if v1 == "frontside" {
			output = "https://i.ibb.co/M2h7RLf/fs-flip.gif"
		}
		if v1 == "backside" {
			output = "https://i.ibb.co/GQHmKMn/bs-flip.gif"
		}
		if v2 == "360" {
			output = "https://i.ibb.co/SXrmc3S/quiztreflip.gif"
		}
		if v2 == "180" {
			output = "https://i.ibb.co/zFz5bv2/fs-flip.gif"
		}
		if v2 == "big" {
			output = "https://i.ibb.co/PmJX75S/big-flip.gif"
		}
		if v2 == "shuv" {
			output = "https://i.ibb.co/C03v02H/quizvarialflip2.gif"
			if v1 == "frontside" {
				output = "https://i.ibb.co/wNgC4qD/hardflip.gif"
			}
		}
	}
	if v3 == "heel" {
		output = "https://i.ibb.co/vwD5b6w/heel.gif"
		if v1 == "frontside" {
			output = "https://i.ibb.co/mbt2YJC/fs-heel.gif"
		}
		if v1 == "backside" {
			output = "https://i.ibb.co/VSPFjgg/bs-heel.gif"
		}
		if v2 == "360" {
			output = "https://i.ibb.co/VM0Z9QG/laserflip.gif"
		}
		if v2 == "180" {
			output = "https://i.ibb.co/PNFhgT4/fs-heel.gif"
		}
		if v2 == "big" {
			output = "https://i.ibb.co/7JY5pJ8/bigheel.gif"
		}
		if v2 == "shuv" {
			output = "https://i.ibb.co/rMpxgxk/inwardheel.gif"
		}
	}
	if v4 == "no comply" && lvl.index > 3 {
		output = "https://i.ibb.co/sq0MSWP/nocomply.gif"
	}
	if v4 == "double" && lvl.index > 3 {
		output = "https://i.ibb.co/b32bDXX/nightmare.gif"
	}
	if v5 == "grab" && lvl.index > 4 {
		output = "https://i.ibb.co/34zQRss/grabby.gif"
	}
	if v5 == "rewind" && lvl.index > 4 {
		output = "https://i.ibb.co/wR4fX2r/ghettobird.gif"
	}
	return output
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
	desc := updateSession(s, m, lvl.index, 0, []int{i0, i1, i2, i3, i4, i5}, set)
	if (i1 == 0 || i1 == 5) && (i2 == 0 || i2 == 5) && (i3 == 0 || i3 == 5) && (i0 == 0 || i0 == 5) && i0+i1+i2+i3 > 9 {
		img := getImg(m, i0, i1, i2, i3, i4, i5, set)
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
				Name:    "SkateDice",
				IconURL: "https://i.ibb.co/jfQPzfm/e8eaa325e2eab7aa995fb25c7bb34f30.jpg",
			},
			Fields: f,
		}
		e.Image = &discordgo.MessageEmbedImage{
			URL: img,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &e)
		return
	}
	if (i1 == 0 || i1 == 5) && (i2 == 0 || i2 == 5) && (i3 == 0 || i3 == 5) && (i0 == 0 || i0 == 5) {
		desc, img, err := giveLetter(s, m, sucky[i4+i5%len(sucky)])
		if err != nil {
			log.Print(err)
			return
		}
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
				Name:    "SkateDice",
				IconURL: "https://i.ibb.co/jfQPzfm/e8eaa325e2eab7aa995fb25c7bb34f30.jpg",
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
	img := getImg(m, i0, i1, i2, i3, i4, i5, set)

	f := []*discordgo.MessageEmbedField{
		{
			Name:  "difficulty",
			Value: lvl.name,
		},
	}
	log.Print(desc)
	if desc == "" {
		return
	}
	e := discordgo.MessageEmbed{
		Color:       lvl.color,
		Title:       fmt.Sprintf("Do a %s.", getTitle(desc)),
		Description: desc,
		URL:         "https://discord.gg/zq3fyV",
		Author: &discordgo.MessageEmbedAuthor{
			URL:     "https://discord.gg/zq3fyV",
			Name:    "SkateDice",
			IconURL: "https://i.ibb.co/jfQPzfm/e8eaa325e2eab7aa995fb25c7bb34f30.jpg",
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

func vertDice(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print("vertdice")
	rollDice(vert, s, m)
}

func jbDice(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print("manualdice")
	rollDice(jb, s, m)
}

func ledDice(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print("ledgedice")
	rollDice(grind, s, m)
}

func reset(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := removeAllLetters(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error resetting the user!")
	}
	delete(sessions, m.Author.ID)
	archiveJSON(os.Getenv("SK8DICE"), &sessions)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@!%s>'s game was reset!", m.Author.ID))
}

func letter(s *discordgo.Session, m *discordgo.MessageCreate) {
	desc, img, err := giveLetter(s, m, "https://i.ibb.co/WPqS8Jw/jarne.gif")
	if err != nil {
		log.Print(err)
		return
	}
	e := discordgo.MessageEmbed{
		Color:       0x000000,
		Title:       "Take the L",
		Description: desc,
		URL:         "https://discord.gg/zq3fyV",
		Author: &discordgo.MessageEmbedAuthor{
			URL:     "https://discord.gg/zq3fyV",
			Name:    "SkateDice",
			IconURL: "https://i.ibb.co/jfQPzfm/e8eaa325e2eab7aa995fb25c7bb34f30.jpg",
		},
	}
	e.Image = &discordgo.MessageEmbedImage{
		URL: img,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &e)
}

func getHandler(s string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	handlers := map[string]func(s *discordgo.Session, m *discordgo.MessageCreate){
		"reg":    regDice,
		"roll":   regDice,
		"vert":   vertDice,
		"bowl":   vertDice,
		"pool":   vertDice,
		"trans":  vertDice,
		"park":   ledDice,
		"ledge":  ledDice,
		"curb":   ledDice,
		"manual": jbDice,
		"malto":  jbDice,
		"jb":     jbDice,
		"reset":  reset,
		"letter": letter,
		"help":   help,
		"list":   help,
		"submit": sub,
		"sub":    sub,
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
