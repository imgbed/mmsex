package apiv2

import (
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type MMv2 struct {
	rds *redis.Client
}

func NewMMv2(rds *redis.Client) *MMv2 {
	return &MMv2{rds: rds}
}

func (obj *MMv2) Alive() int64 {
	return 60
}

func (obj *MMv2) Description() string {
	return "美女提神图片"
}

func (ojb *MMv2) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid, err := ojb.rds.SRandMember(store.SetImg).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2, _ := ojb.rds.HGet(store.HashImg, imgid).Result()

	if img2 == "" {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, "暂无图片")
	} else {
		tc = tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{
			tgbotapi.NewInputMediaPhoto(img2),
		})
	}
	return
}
