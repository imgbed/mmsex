package main

import (
	"github.com/gohouse/tgbot"
	"github.com/gohouse/tgbot/api"
)

func main() {
	bot := tgbot.NewTgBot(
		//tgbot.SetAliveTime(999),
		)
	//bot.Use(cors.UserImg)

	//bot.Listen("asian", apiv3.NewAsian(bot.Store))
	//bot.Listen("kfc", apiv3.NewKfc(bot.Store))
	//bot.Listen("mm", apiv2.NewMMv2(bot.Store))
	//bot.Listen("neihan", apiv3.NewNeihan(bot.Store))
	////bot.Listen("duanzi", apiv3.NewDuanzi(bot.Store))
	////bot.Listen("cc", apiv3.NewCC(bot.Store))
	bot.Listen("djt", api.NewPCS())
	//bot.Listen("xq", apiv3.NewXiongQi(bot.Store))
	//bot.Listen("xqs", apiv3.NewXiongQiSex(bot.Store))
	//bot.Listen("sfz", apiv3.NewSfz(bot.Store))
	//bot.Listen("jingpin", apiv3.NewJingPin(bot.Store))
	//bot.Listen("video", apiv3.NewVideo(bot.Store))
	//bot.Listen("img", bak.NewUserImg(bak.UserImgImg,bot.Store))
	//bot.Listen("gif", bak.NewUserImg(bak.UserImgGif,bot.Store))

	bot.Run()
}
