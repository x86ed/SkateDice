package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var lucky = []string{"https://i.ibb.co/ss7krZ0/900.gif", "https://i.ibb.co/RjL95p8/lyon.gif", "https://i.ibb.co/K5hWG5S/sunset.gif", "https://i.ibb.co/9rb36tf/booth.gif", "https://i.ibb.co/1bR0b1w/mo.gif", "https://i.ibb.co/9r0ttwn/lakai.gif", "https://i.ibb.co/yFn8Txh/wall.gif", "https://i.ibb.co/HGwyRHH/tree.gif", "https://i.ibb.co/NScxfNR/carlsbad.gif", "https://i.ibb.co/cYSqN0K/staygold.gif", "https://i.ibb.co/dJ4pwdB/kirchart.gif"}
var sucky = []string{"https://i.ibb.co/WPqS8Jw/jarne.gif", "https://i.ibb.co/6J1KYt4/skin.png", "https://i.ibb.co/HLy99YM/duffman.gif", "https://i.ibb.co/VpK322W/wasted.gif", "https://i.ibb.co/zmM4K9m/boulala.gif", "https://i.ibb.co/YNPQPS2/sucky.gif", "https://i.ibb.co/xS3Pt94/sacked.gif", "https://i.ibb.co/qdRdypn/fence.gif", "https://i.ibb.co/b1vdd94/carded.gif", "https://i.ibb.co/znDb8d9/2xZZG1.gif", "https://i.ibb.co/QQzh97P/Pessimistic-Brilliant-Hammerheadbird-size-restricted.gif", "https://i.ibb.co/HTwv8gf/xnGG1r.gif", "https://i.ibb.co/nR0vnHZ/Embellished-Careless-Hare-small.gif", "https://i.ibb.co/RYgPZF9/P7kLkW.gif"}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func containsVal(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func archiveJSON(fn string, ty interface{}) {
	f, err := os.Create(fn)
	if err != nil {
		return
	}

	defer f.Close()

	arch, err := json.Marshal(ty)
	if err != nil {
		return
	}
	c, err := f.Write(arch)
	if err != nil {
		return
	}
	fmt.Println("bytes: ", c)
}

func unarchiveJSON(fn string, ty interface{}) {
	if fileExists(fn) {
		dat, err := ioutil.ReadFile(fn)
		if err != nil {
			return
		}
		json.Unmarshal(dat, ty)
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func hasRole(m *discordgo.MessageCreate, id string) bool {
	fmt.Printf("roles %+v\n", m.Message.Member.Roles)
	if m.Message.Member != nil {
		if containsVal(m.Message.Member.Roles, id) > -1 {
			return true
		}
	}
	return false
}

func getLevel(m *discordgo.MessageCreate) int {
	if hasRole(m, "795427084540968980") {
		return 4
	}
	if hasRole(m, "795426715056603157") {
		return 3
	}
	if hasRole(m, "795426286838218773") {
		return 2
	}
	if hasRole(m, "795426137248366603") {
		return 1
	}
	if hasRole(m, "795425922920480779") {
		return 0
	}
	return -1
}

func getDif(m *discordgo.MessageCreate) difficulty {
	if hasRole(m, "795528642151055400") {
		return difficulty{0xFF00FF, "Yummy", 5}
	}
	if hasRole(m, "795528892613525544") {
		return difficulty{0xFF0000, "Pro", 4}
	}
	if hasRole(m, "795529016831377468") {
		return difficulty{0xFFFF00, "Am", 3}
	}
	if hasRole(m, "795529229613269022") {
		return difficulty{0x00FF00, "Flow", 2}
	}
	return difficulty{0x666666, "Grom", 1}
}

func removeAllLetters(s *discordgo.Session, m *discordgo.MessageCreate) error {
	lvls := []string{
		"795425922920480779",
		"795426137248366603",
		"795426286838218773",
		"795426715056603157",
		"795427084540968980",
		"795528642151055400",
		"795528892613525544",
		"795529016831377468",
		"795529229613269022",
	}
	for _, v := range lvls {
		err := s.GuildMemberRoleRemove(m.GuildID, m.Author.ID, v)
		if err != nil {
			log.Print("err removing level ", v)
			//return err
		}
	}
	return nil
}

func addLevel(s *discordgo.Session, m *discordgo.MessageCreate, r string) error {
	log.Print(r)
	err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, r)
	if err != nil {
		log.Print("add lev err")
		return err
	}
	delete(sessions, m.Author.ID)
	archiveJSON(os.Getenv("SK8DICE"), &sessions)
	return nil
}

func addDif(s *discordgo.Session, mA *discordgo.MessageReactionAdd) {
	mem, err := s.GuildMember(mA.GuildID, mA.UserID)
	mm, err := s.ChannelMessage(mA.ChannelID, mA.MessageID)
	if err != nil {
		return
	}
	mm.Member = mem
	mm.GuildID = mA.GuildID
	m := &discordgo.MessageCreate{Message: mm}
	d := getDif(m)
	if d.index >= 5 {
		return
	}
	next := []string{
		"",
		"",
		"795529229613269022",
		"795529016831377468",
		"795528892613525544",
		"795528642151055400",
	}
	// refactor
	err = addLevel(s, m, next[d.index+1])
	if err != nil {
		log.Print("add level error")
		return
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@!%s> leveled up to <@&%s>!", mem.User.ID, next[d.index+1]))
}

func giveLetter(s *discordgo.Session, m *discordgo.MessageCreate, img string) (string, string, error) {
	lvl := getLevel(m)
	log.Printf("level letter: %d", lvl)
	lvl++
	log.Printf("level letter: %d", lvl)
	if lvl >= 4 {
		err := removeAllLetters(s, m)
		if err != nil {
			log.Print("couldn't remove letters")
			log.Print(err)
		}
		delete(sessions, m.Author.ID)
		archiveJSON(os.Getenv("SK8DICE"), &sessions)
		return "You lose! Don't @me " + m.Message.Author.Username, "https://i.ibb.co/31bn5D8/SCRAPELINE-SMD-EPISODE-6.gif", err
	}
	lvls := []string{
		"795425922920480779",
		"795426137248366603",
		"795426286838218773",
		"795426715056603157",
		"795427084540968980",
	}
	ltr := []string{
		"S",
		"K",
		"A",
		"T",
		"E",
	}
	err := addLevel(s, m, lvls[lvl])
	if err != nil {
		log.Print("give Letter err")
		return "", "", err
	}
	return "Oof. you got a(n)" + ltr[lvl] + "!", img, nil

}

func giveLetterR(s *discordgo.Session, mA *discordgo.MessageReactionAdd, img string) (string, string, error) {
	fmt.Printf(" sesh: %+v\n ma: %+v\n", s, mA)
	mem, err := s.GuildMember(mA.GuildID, mA.UserID)
	mm, err := s.ChannelMessage(mA.ChannelID, mA.MessageID)
	if err != nil {
		return "", "", err
	}
	mm.Member = mem
	mm.GuildID = mA.GuildID
	m := &discordgo.MessageCreate{Message: mm}
	lvl := getLevel(m)
	log.Printf("level: %d", lvl)
	lvl++
	if lvl >= 4 {
		err := removeAllLetters(s, m)
		if err != nil {
			log.Print(err)
		}
		delete(sessions, m.Author.ID)
		archiveJSON(os.Getenv("SK8DICE"), &sessions)
		return "You lose! Don't @me " + m.Message.Author.Username, "https://i.ibb.co/31bn5D8/SCRAPELINE-SMD-EPISODE-6.gif", err
	}
	lvls := []string{
		"795425922920480779",
		"795426137248366603",
		"795426286838218773",
		"795426715056603157",
		"795427084540968980",
	}
	ltr := []string{
		"S",
		"K",
		"A",
		"T",
		"E",
	}
	err = addLevel(s, m, lvls[lvl])
	if err != nil {
		return "", "", err
	}
	return "Oof. you got a(n)" + ltr[lvl] + "!", img, nil

}
