package bak

import "github.com/gohouse/golib/tgbot/adapter"

type MM adapter.LocalPhoto

func NewMM(cooldownduration ...int) adapter.ImgAdapter {
	var cooldown int
	if len(cooldownduration)>0 {
		cooldown = cooldownduration[0]
	}
	return adapter.NewLocalPhoto("/Users/kevin/Downloads/mm_img/","妹妹图",cooldown)
}
