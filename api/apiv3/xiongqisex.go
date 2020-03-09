package apiv3

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type XiongQiSex struct {
	rds *redis.Client
}

func NewXiongQiSex(rds *redis.Client) *XiongQiSex {
	return &XiongQiSex{rds: rds}
}

func (obj *XiongQiSex) Alive() int64 {
	return 30
}

func (obj *XiongQiSex) Description() string {
	return fmt.Sprintf("露点大凶器,时效:%vs",obj.Alive())
}

func (ojb *XiongQiSex) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid, err := ojb.rds.SRandMember(store.SetXiongQiSex).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2, err := ojb.rds.HGet(store.HashXiongQiSex, imgid).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}

	if img2 == "" {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, "暂无图片")
	} else {
		tc = tgbotapi.NewPhotoShare(update.Message.Chat.ID, img2)
		//tgbotapi.NewMediaGroup(0,[]interface{}{
		//	tgbotapi.NewPhotoShare(update.Message.Chat.ID, img2),
		//})
	}
	return
}
