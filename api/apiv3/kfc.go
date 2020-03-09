package apiv3

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type Kfc struct {
	rds *redis.Client
}

func NewKfc(rds *redis.Client) *Kfc {
	return &Kfc{rds: rds}
}

func (obj *Kfc) Alive() int64 {
	return 15
}

func (obj *Kfc) Description() string {
	return fmt.Sprintf("你懂的,欧美货(特殊素材,时效 %vs)",obj.Alive())
}

func (ojb *Kfc) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid, err := ojb.rds.SRandMember(store.SetKfc).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2, err := ojb.rds.HGet(store.HashKfc, imgid).Result()
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
