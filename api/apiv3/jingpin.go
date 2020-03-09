package apiv3

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type JingPin struct {
	rds *redis.Client
}

func NewJingPin(rds *redis.Client) *JingPin {
	return &JingPin{rds: rds}
}

func (obj *JingPin) Alive() int64 {
	return 15
}

func (obj *JingPin) Description() string {
	return fmt.Sprintf("精品图(上品好货,时效 %vs)",obj.Alive())
}

func (ojb *JingPin) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid, err := ojb.rds.SRandMember(store.SetJingPin).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2, err := ojb.rds.HGet(store.HashJingPin, imgid).Result()
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
