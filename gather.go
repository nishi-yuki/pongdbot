package pongdbot

import (
	"os"
	"strings"

	"github.com/nishi-yuki/pongdbot/taskrunner"
)

const (
	token    = "DISCORD_TOKEN"
	userlist = "USER_LIST"
	prefix   = "PREFIX"
)

func GatherEnvVar(t *taskrunner.TaskRunner) (bot *Bot) {
	token := os.Getenv(token)
	userlist := strings.Fields(os.Getenv(userlist))
	prefix := os.Getenv(prefix)
	if prefix == "" {
		prefix = "."
	}
	bot = New(token, userlist, prefix, t)
	return
}
