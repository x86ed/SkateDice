package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// // Create a new Discord session using the provided bot token.
	// if fileExists(os.Getenv("SK8DICE")) {
	// 	dat, err := ioutil.ReadFile(os.Getenv("SK8DICE"))
	// 	if err != nil {
	// 		return
	// 	}
	// 	json.Unmarshal(dat, &cases)
	// }

	dg, err := discordgo.New("Bot " + os.Getenv("DICEID"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(Stickman)
	dg.AddHandler(Reaction)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
