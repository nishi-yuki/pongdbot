package pongdbot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/nishi-yuki/pongdbot/taskrunner"
)

type Bot struct {
	authToken  string
	UserList   UserList
	Prefix     string
	TaskRunner *taskrunner.TaskRunner
}

func New(authToken string, userlist UserList,
	prefix string, t *taskrunner.TaskRunner) (bot *Bot) {
	bot = &Bot{
		authToken:  authToken,
		UserList:   userlist,
		Prefix:     prefix,
		TaskRunner: t,
	}
	return
}

func (bot *Bot) Start() {
	discord, err := discordgo.New("Bot " + bot.authToken)
	if err != nil {
		log.Fatal(err)
		return
	}

	discord.AddHandler(bot.handler)
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start Bot")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	fmt.Println("exit")
	discord.Close()
}

func (bot *Bot) handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, bot.Prefix) {
		return
	}
	if !bot.UserList.Contains(m.Author) {
		fmt.Println(m.Author, "is not authorized.")
		return
	}

	msgl := strings.Fields(m.Content)
	taskName := strings.TrimPrefix(msgl[0], bot.Prefix)

	fmt.Print("Run \"" + taskName + "\" ")
	res, err := bot.TaskRunner.Run(taskName, m, msgl[1:])
	if err != nil {
		errMsg := "Error: Task \"" + taskName + "\" is not defined."
		s.ChannelMessageSend(m.ChannelID, errMsg)
		fmt.Println("is not defined.")
		return
	}
	fmt.Println("ok.")
	s.ChannelMessageSend(m.ChannelID, res)
}
