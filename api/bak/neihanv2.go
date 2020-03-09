package bak

import (
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
	"strings"
)

type Neihanv2 struct {
	rds *redis.Client
}

func NewNeihanv2(rds *redis.Client) *Neihanv2 {
	return &Neihanv2{rds:rds}
}

func (obj *Neihanv2) Description() string {
	return "内涵gif动图"
}

func (ojb *Neihanv2) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	imgid,err := ojb.rds.SRandMember(store.SetGif).Result()
	if err!=nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	img2,_ := ojb.rds.HGet(store.HashGif,imgid).Result()

	if img2 == ""{
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, "暂无图片")
	}else {
		// https://cdn.jsdelivr.net/gh/imgbed/images@
		var gh = "https://raw.githubusercontent.com/imgbed/images/"
		var cdnpre = "https://cdn.jsdelivr.net/gh/imgbed/images@"

		//tc = tgbotapi.NewMediaGroup(update.Message.Chat.ID,[]interface{}{
		//	//tgbotapi.NewInputMediaAnimation(img2),
		//	tgbotapi.NewInputMediaVideo(strings.Replace(img2, gh,cdnpre,-1)),
		//	//tgbotapi.NewInputMediaAnimation(strings.Replace(img2, gh,cdnpre,-1)),
		//})
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, strings.Replace(img2, gh,cdnpre,-1))
		//tc = tgbotapi.NewAnimationUpload(update.Message.Chat.ID, img2)
	}
	return
}
