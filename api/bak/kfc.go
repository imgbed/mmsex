package bak

import "github.com/gohouse/golib/tgbot/adapter"

type KFC adapter.LocalMedia

func NewKFC(cooldownduration ...int) adapter.ImgAdapter {
	var cooldown int
	if len(cooldownduration)>0 {
		cooldown = cooldownduration[0]
	}
	return adapter.NewLocalMedia("/Users/kevin/Downloads/img/gif/kfc/","你懂的(文件较大,加载时间很长)",cooldown)
}
