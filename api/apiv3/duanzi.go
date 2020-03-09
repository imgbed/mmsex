package apiv3

import (
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/curl"
	"github.com/gohouse/golib/t"
)

type Duanzi struct {
	rds *redis.Client
}

func NewDuanzi(rds *redis.Client) *Duanzi {
	return &Duanzi{rds: rds}
}

func (obj *Duanzi) Alive() int64 {
	return 600
}

func (obj *Duanzi) Description() string {
	return "段子"
}

func (ojb *Duanzi) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	res,err :=curl.Get("https://api.apiopen.top/getJoke?type=text&count=1")
	if err!=nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	res2 := t.New(res.Extract("result")).SliceMapString()

	if len(res2)==0 {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, "获取段子失败,请重试")
	} else {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, res2[0]["text"].String())
	}
	return
}
