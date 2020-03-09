package bak

import (
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/store"
)

type UserImgType int

const (
	UserImgImg UserImgType = iota
	UserImgGif
)

type UserImg struct {
	imgs   []string
	gifs   []string
	adtype UserImgType
	note   string
	store  *redis.Client
}

func NewUserImg(adtype UserImgType, store *redis.Client) *UserImg {
	ad := &UserImg{adtype: adtype, store: store}

	return ad
}

func (obj *UserImg) Alive() int64 {
	return 60
}

func (mm *UserImg) Description() string {
	if mm.adtype == UserImgGif {
		return "自动存储用户发送的动态gif图复用"
	}
	return "自动存储用户发送的静态图复用"
}

func (mm *UserImg) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	if mm.adtype == UserImgGif {
		imgid, err := mm.store.SRandMember(store.SetUserGif).Result()
		if err != nil {
			tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
			return
		}
		img2, _ := mm.store.HGet(store.HashUserGif, imgid).Result()

		if img2 == "" {
			tc = tgbotapi.NewMessage(update.Message.Chat.ID, "暂无图片")
		} else {
			tc = tgbotapi.NewAnimationShare(update.Message.Chat.ID, img2)
		}
	} else if mm.adtype == UserImgImg {
		imgid, err := mm.store.SRandMember(store.SetUserImg).Result()
		if err != nil {
			tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
			return
		}
		img2, _ := mm.store.HGet(store.HashUserImg, imgid).Result()
		if img2 == "" {
			tc = tgbotapi.NewMessage(update.Message.Chat.ID, "暂无图片")
		} else {
			tc = tgbotapi.NewPhotoShare(update.Message.Chat.ID, img2)
		}
	}
	return
}
