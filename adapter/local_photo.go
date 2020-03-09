package adapter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot/util"
)

type LocalPhoto struct {
	*ImgAlive
	localPath string
	note      string // 说明
	cooldown  int
}

func NewLocalPhoto(localPath, note string, cooldown int) *LocalPhoto {
	return &LocalPhoto{localPath: localPath, note: note, cooldown: cooldown}
}

func (mm *LocalPhoto) Description() string {
	return mm.note
}

func (mm *LocalPhoto) GetCooldown() int {
	return mm.cooldown
}

func (mm *LocalPhoto) BuildChattable(update *tgbotapi.Update, cmdArgs string) tgbotapi.Chattable {
	img := util.GetRandFileFromDir(mm.localPath)
	return tgbotapi.NewPhotoUpload(update.Message.Chat.ID, img)
}
