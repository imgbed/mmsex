package adapter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/util"
)

type LocalMedia struct {
	*ImgAlive
	localPath string
	note      string // 说明
	cooldown  int
}

func NewLocalMedia(localPath, note string, cooldown int) *LocalMedia {
	return &LocalMedia{localPath: localPath, note: note, cooldown: cooldown}
}

func (mm *LocalMedia) Description() string {
	return mm.note
}

func (mm *LocalMedia) GetCooldown() int {
	return mm.cooldown
}

func (mm *LocalMedia) BuildChattable(update *tgbotapi.Update, cmdArgs string) tgbotapi.Chattable {
	img := util.GetRandFileFromDir(mm.localPath)
	return tgbotapi.NewAnimationUpload(update.Message.Chat.ID, img)
	//return tgbotapi.NewDocumentUpload(update.Message.Chat.ID, img)
}
