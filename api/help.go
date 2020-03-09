package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Help struct {
	cooldown int
}

func NewHelp() *Help {
	return &Help{}
}

func (obj *Help) Alive() int64 {
	return 60
}

func (mm *Help) Description() string {
	return "帮助"
}

func (mm *Help) GetCooldown() int {
	return mm.cooldown
}

func (mm *Help) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	return
}
