package adapter

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type ImgAdapter interface {
	Description() string
	BuildChattable(update *tgbotapi.Update,cmdArgs string) tgbotapi.Chattable
	//After(msg *tgbotapi.Message)
	Alive() int64
}

type ImgAlive struct {
	alive int64
}

func NewImgAlive(a int64) *ImgAlive {
	return &ImgAlive{a}
}

func (ia *ImgAlive) Alive() int64 {
	return ia.alive
}

type PersitAdapter interface {
	Store(arg interface{}) error
	Load(arg interface{}) error
}