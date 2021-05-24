package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nishi-yuki/pongdbot"
	"github.com/nishi-yuki/pongdbot/taskrunner"
)

func main() {
	tr := taskrunner.New()
	tr.Add("ping", ping)
	bot := pongdbot.GatherEnvVar(tr)
	bot.Start()
}

func ping(m *discordgo.MessageCreate, _ []string) string {
	return "Pong! " + m.Author.String()
}
