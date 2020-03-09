package bak

import "github.com/gohouse/golib/tgbot/adapter"

type Neihan adapter.LocalPhoto

func NewNeihan(cooldownduration ...int) adapter.ImgAdapter {
	var cooldown int
	if len(cooldownduration)>0 {
		cooldown = cooldownduration[0]
	}
	return adapter.NewLocalMedia("/Users/kevin/Downloads/img/gif/neihan/","gif-内涵图(图片较大,加载需要10s以上)",cooldown)
}
