package apiv3

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type Sfz struct {
	rds *redis.Client
}

func NewSfz(rds *redis.Client) *Sfz {
	return &Sfz{rds: rds}
}

func (obj *Sfz) Alive() int64 {
	return 30
}

func (obj *Sfz) Description() string {
	return fmt.Sprintf("死肥宅福利图")
}

func (ojb *Sfz) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid, err := ojb.rds.SRandMember(store.SetSfz).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2, err := ojb.rds.HGet(store.HashSfz, imgid).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}

	if img2 == "" {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, "暂无图片")
	} else {
		tc = tgbotapi.NewPhotoShare(update.Message.Chat.ID, img2)
	}
	return
}
