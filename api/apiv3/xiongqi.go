package apiv3

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type XiongQi struct {
	rds *redis.Client
}

func NewXiongQi(rds *redis.Client) *XiongQi {
	return &XiongQi{rds: rds}
}

func (obj *XiongQi) Alive() int64 {
	return 60
}

func (obj *XiongQi) Description() string {
	return fmt.Sprintf("G罩杯大凶器,时效:%vs",obj.Alive())
}

func (ojb *XiongQi) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid, err := ojb.rds.SRandMember(store.SetXiongQi).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2, err := ojb.rds.HGet(store.HashXiongQi, imgid).Result()
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
