package apiv3

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/curl"
)

type CC struct {
	rds *redis.Client
}

func NewCC(rds *redis.Client) *CC {
	return &CC{rds: rds}
}

func (obj *CC) Alive() int64 {
	return 180
}

func (obj *CC) Description() string {
	return fmt.Sprintf("聊天机器人,使用 /cc {msg}")
}

func (ojb *CC) BuildChattable(update *tgbotapi.Update, cmdArgs string) (tc tgbotapi.Chattable) {
	resp,err := curl.Get(fmt.Sprintf("http://api.qingyunke.com/api.php?key=free&appid=0&msg=%s",cmdArgs))
	if err!=nil {
		tc = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		return
	}
	tc = tgbotapi.NewMessage(update.Message.Chat.ID, resp.Extract("content").String())
	return
}
