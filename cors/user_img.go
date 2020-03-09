package cors

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/tgbot"
	"github.com/gohouse/golib/tgbot/store"
)

func UserImg(c *tgbot.Context) {
	// 自动截留用户发送的图片
	// 记录动图信息
	if c.Update.Message.Animation != nil {
		switch c.Update.Message.Chat.ID {

		case -370863718: // 删除图	-370863718
			// 获取要删除图片的 id
			c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID,
				fmt.Sprintf("销毁 %s 发送的图片,id为: %s",
					c.Update.Message.From.FirstName,
					c.Update.Message.Animation.FileID)))
			//c.Update.Message.Animation
			//c.TgBot.Store.SRem(store.SetUserGif, c.Update.Message.Animation.FileUniqueID)
			//c.TgBot.Store.HDel(store.HashUserGif, c.Update.Message.Animation.FileUniqueID)
			c.TgBot.Store.SRem(annimationChatId[-370863718]["set"], c.Update.Message.Animation.FileUniqueID)
			c.TgBot.Store.HDel(annimationChatId[-370863718]["hash"], c.Update.Message.Animation.FileUniqueID)
		default: // 保存图
			if r, ok := annimationChatId[c.Update.Message.Chat.ID]; ok {
				c.TgBot.Store.SAdd(r["set"], c.Update.Message.Animation.FileUniqueID)
				c.TgBot.Store.HSet(r["hash"], c.Update.Message.Animation.FileUniqueID, c.Update.Message.Animation.FileID)
				c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID, fmt.Sprintf("已收录 %s 发送的新动图,可通过 %s 随机调用", c.Update.Message.From.FirstName, r["cmd"])))
			}
		}
		c.Abort()
	} else if c.Update.Message.Video != nil {
		switch c.Update.Message.Chat.ID {
		//case -370863718: // 删除图	-370863718
		//	// 获取要删除图片的 id
		//	c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID,
		//		fmt.Sprintf("销毁 %s 发送的图片,id为: %s",
		//			c.Update.Message.From.FirstName,
		//			c.Update.Message.Animation.FileID)))
		//	//c.Update.Message.Animation
		//	//c.TgBot.Store.SRem(store.SetUserGif, c.Update.Message.Animation.FileUniqueID)
		//	//c.TgBot.Store.HDel(store.HashUserGif, c.Update.Message.Animation.FileUniqueID)
		//	c.TgBot.Store.SRem(annimationChatId[-370863718]["set"], c.Update.Message.Animation.FileUniqueID)
		//	c.TgBot.Store.HDel(annimationChatId[-370863718]["hash"], c.Update.Message.Animation.FileUniqueID)
		default: // 保存视频
			if r, ok := annimationChatId[c.Update.Message.Chat.ID]; ok {
				fmt.Println("r:", r)
				fmt.Println("FileUniqueID:", c.Update.Message.Video.FileUniqueID)
				fmt.Println("FileID:", c.Update.Message.Video.FileID)
				c.TgBot.Store.SAdd(r["set"], c.Update.Message.Video.FileUniqueID)
				c.TgBot.Store.HSet(r["hash"], c.Update.Message.Video.FileUniqueID, c.Update.Message.Video.FileID)
				c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID, fmt.Sprintf("已收录 %s 发送的新视频,可通过 %s 随机调用", c.Update.Message.From.FirstName, r["cmd"])))
			}
		}
		c.Abort()
	} else if len(c.Update.Message.Photo) > 0 { // 记录图片信息
		var idx = len(c.Update.Message.Photo) - 1
		var imgid = (c.Update.Message.Photo)[idx]
		// 本地删除
		// 删除图	-370863718 :1022401312,
		var chatId = c.Update.Message.Chat.ID
		switch chatId {
		case -370863718: // del_img 群
			// 获取要删除图片的 id
			c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID,
				fmt.Sprintf("销毁 %s 发送的图片,id为: %s",
					c.Update.Message.From.FirstName,
					imgid.FileUniqueID)))
			c.TgBot.Store.SRem(store.SetUserImg, imgid.FileUniqueID)
			c.TgBot.Store.HDel(store.HashUserImg, imgid.FileUniqueID)
		default:
			if r, ok := photoChatId[c.Update.Message.Chat.ID]; ok {
				// 本地存储
				c.TgBot.Store.SAdd(r["set"], imgid.FileUniqueID)
				c.TgBot.Store.HSet(r["hash"], imgid.FileUniqueID, imgid.FileID)
				c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID, fmt.Sprintf("已收录 %s 发送的新图片,可通过 %s 随机调用", c.Update.Message.From.FirstName, r["cmd"])))

			} else if chatId == 1022401312 { // doutu_bot 私聊,储存图片
				// 本地存储
				c.TgBot.Store.SAdd(store.SetUserImg, imgid.FileUniqueID)
				c.TgBot.Store.HSet(store.HashUserImg, imgid.FileUniqueID, imgid.FileID)

				c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID, fmt.Sprintf("已收录 %s 发送的新图片,可通过 /img 随机调用", c.Update.Message.From.FirstName)))
			}
		}
		//if chatId == -370863718 { // del_img 群
		//	// 获取要删除图片的 id
		//	c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID,
		//		fmt.Sprintf("销毁 %s 发送的图片,id为: %s",
		//			c.Update.Message.From.FirstName,
		//			imgid.FileUniqueID)))
		//	c.TgBot.Store.SRem(store.SetUserImg, imgid.FileUniqueID)
		//	c.TgBot.Store.HDel(store.HashUserImg, imgid.FileUniqueID)
		//} else if chatId == 1022401312 { // doutu_bot 私聊,储存图片
		//	// 本地存储
		//	c.TgBot.Store.SAdd(store.SetUserImg, imgid.FileUniqueID)
		//	c.TgBot.Store.HSet(store.HashUserImg, imgid.FileUniqueID, imgid.FileID)
		//
		//	c.TgBot.BotAPI.Send(tgbotapi.NewMessage(c.Update.Message.Chat.ID, fmt.Sprintf("已收录 %s 发送的新图片,可通过 /img 随机调用", c.Update.Message.From.FirstName)))
		//}
		c.Abort()
	} else {
		c.Next()
	}
}
