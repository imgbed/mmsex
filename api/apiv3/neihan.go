package apiv3

import (
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type Neihan struct {

	rds *redis.Client
}

func NewNeihan(rds *redis.Client) *Neihan {
	return &Neihan{rds: rds}
}

func (obj *Neihan) Alive() int64 {
	return 60
}

func (obj *Neihan) Description() string {
	return "内涵gif动图"
}

func (ojb *Neihan) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid, err := ojb.rds.SRandMember(store.SetNeihan).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2, err := ojb.rds.HGet(store.HashNeihan, imgid).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}

	if img2 == "" {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, "暂无图片")
	} else {
		tc = tgbotapi.NewAnimationShare(update.Message.Chat.ID, img2)
	}
	return
}
