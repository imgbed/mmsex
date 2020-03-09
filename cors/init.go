package cors

import "github.com/gohouse/golib/tgbot/store"

var annimationChatId = map[int64]map[string]string {
	// 一般用户发的gif表情
	-370863718: {"set":store.SetUserGif,"hash":store.HashUserGif,"cmd":"/gif"},
	//-351168855: {"set":store.SetKfc,"hash":store.HashKfc,"cmd":"/kfc"},
	// 一般kfc sex动图
	-351168855: {"set":store.SetKfc,"hash":store.HashKfc,"cmd":"/kfc"},
	// 内涵而不露点动图  -1001434990250
	-1001434990250: {"set":store.SetNeihan,"hash":store.HashNeihan,"cmd":"/neihan"},
	// img_收集精品动图
	-364313246: {"set":store.SetJingPin,"hash":store.HashJingPin,"cmd":"/jingpin"},
	// img_视频收集
	//-345806495: {"set":store.SetVideo,"hash":store.HashVideo,"cmd":"/video"},
	-1001261221279: {"set":store.SetVideo,"hash":store.HashVideo,"cmd":"/video"},
}

var photoChatId = map[int64]map[string]string{
	// img_死肥宅静态图片收集
	-1001223946282: {"set":store.SetSfz,"hash":store.HashSfz,"cmd":"/sfz"},
}
