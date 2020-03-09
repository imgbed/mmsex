package api

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Help struct {
}

func NewHelp() *Help {
	return &Help{}
}

func (obj *Help) Alive() int64 {
	return 0
}

func (obj *Help) Description() string {
	return fmt.Sprintf("帮助")
}

func (ojb *Help) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	tc = tgbotapi.NewMessage(update.Message.Chat.ID, "Help()")
	return
}