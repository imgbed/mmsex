package cors

import (
	"github.com/gohouse/tgbot"
)

func Sex(c *tgbot.Context) {
	// 如果是私聊, 则保存图片,动图,视频到指定命令
	if c.Update.Message.Chat.ID == int64(c.Update.Message.From.ID) {
		// gif,img,video
	} else {
		// 如果是群聊, 则保存图片,动图,视频到另外的命令
		// autogif
	}

	// 自动截留用户发送的图片
	// 记录动图信息
	//c.TgBot.Store.SAdd(r["set"], c.Update.Message.Animation.FileUniqueID)
	//c.TgBot.Store.HSet(r["hash"], c.Update.Message.Animation.FileUniqueID, c.Update.Message.Animation.FileID)
	//c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID, fmt.Sprintf("已收录 %s 发送的新动图,可通过 %s 随机调用", c.Update.Message.From.FirstName, r["cmd"])))
	c.Abort()
}
