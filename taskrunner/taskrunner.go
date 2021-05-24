package taskrunner

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

type Task func(m *discordgo.MessageCreate, args []string) string

type TaskRunner struct {
	taskMap map[string]Task
}

func New() (t *TaskRunner) {
	t = &TaskRunner{make(map[string]Task)}
	return
}

func (tm *TaskRunner) Add(name string, t Task) {
	tm.taskMap[name] = t
}

func (tm TaskRunner) Run(name string, m *discordgo.MessageCreate, args []string) (res string, err error) {
	task := tm.taskMap[name]
	if task == nil {
		err = errors.New("task not found")
		return
	}
	res = task(m, args)
	err = nil
	return
}
