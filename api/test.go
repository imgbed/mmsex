package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type Test struct {
	cooldown int
}

func NewTest() *Test {
	return &Test{}
}

func (mm *Test) Description() string {
	return "test"
}

func (mm *Test) GetCooldown() int {
	return mm.cooldown
}

func (mm *Test) BuildChattable(update *tgbotapi.Update,cmdArgs string) (tc tgbotapi.Chattable) {
	tc = tgbotapi.NewMessage(update.Message.Chat.ID, "test text.")
	//tc = tgbotapi.NewPhotoShare(update.Message.Chat.ID,
	//	"AQADJ4bfMgAECTcAAg")
	if cmdArgs=="set" {
		store.Redis().Set("a",1,0)
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, "set a=1")
	}
	if cmdArgs=="get" {
		r,_:=store.Redis().Get("a").Result()
		tc = tgbotapi.NewMessage(update.Message.Chat.ID,
			"get a result is: "+r)
	}

	return
}
