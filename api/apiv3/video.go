package apiv3

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type Video struct {
	rds *redis.Client
}

func NewVideo(rds *redis.Client) *Video {
	return &Video{rds: rds}
}

func (obj *Video) Alive() int64 {
	return 10
}

func (obj *Video) Description() string {
	return fmt.Sprint("视频盛宴")
}

func (ojb *Video) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid, err := ojb.rds.SRandMember(store.SetVideo).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2, err := ojb.rds.HGet(store.HashVideo, imgid).Result()
	if err != nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}

	if img2 == "" {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, "暂无视频")
	} else {
		tc = tgbotapi.NewVideoShare(update.Message.Chat.ID, img2)
	}
	return
}
