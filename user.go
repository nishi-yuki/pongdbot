package pongdbot

import (
	"github.com/bwmarrin/discordgo"
)

type UserList []string

func (ul *UserList) Contains(du *discordgo.User) bool {
	for _, u := range *ul {
		if u == du.String() {
			return true
		}
	}
	return false
}
