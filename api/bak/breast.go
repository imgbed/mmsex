package bak

import "github.com/gohouse/golib/tgbot/adapter"

type Breast adapter.LocalPhoto

func NewBreast(cooldownduration ...int) adapter.ImgAdapter {
	var cooldown int
	if len(cooldownduration)>0 {
		cooldown = cooldownduration[0]
	}
	return adapter.NewLocalMedia("/Users/kevin/Downloads/img/gif/breast/","gif-胸大无脑",cooldown)
}
