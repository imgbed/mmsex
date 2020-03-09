package store

import "github.com/go-redis/redis/v7"

const (
	SetGif  = "s:gif"	// github的动图
	SetImg  = "s:img"	// github的mm
	//SetUserImg  = "s:xiongqi"	// 用户发的静,doutu_bot 私聊  userimg
	SetUserImg  = "s:userimg"	// 用户发的静,doutu_bot 私聊  userimg
	SetUserGif  = "s:usergif"	// 用户发的动图
	SetNeihan  = "s:neihan"	// chat的 neihan
	SetDuanzi  = "s:duanzi"	// duanzi
	SetKfc  = "s:kfc"	// chat的 kfc
	SetAsian  = "s:asian"	// chat的 asian
	SetJingPin  = "s:jingpin"	// chat的 精品
	SetVideo  = "s:video"	// chat的 video
	SetSfz  = "s:sfz"	// chat的 sfz
	SetXiongQi = "s:xiongqi"	// 胸器图
	SetXiongQiSex = "s:xiongqisex"	// 胸器色图

	HashGif = "h:gif"
	HashImg = "h:img"
	//HashUserImg = "h:xiongqi"
	HashUserImg = "h:userimg"
	HashUserGif = "h:usergif"
	HashNeihan = "h:neihan"
	HashDuanzi = "h:duanzi"
	HashKfc = "h:kfc"
	HashAsian = "h:asian"
	HashJingPin = "h:jingpin"
	HashVideo = "h:video"
	HashSfz = "h:sfz"
	HashXiongQi = "h:xiongqi"	// 胸器图
	HashXiongQiSex = "h:xiongqisex"	// 胸器色图
)
var client *redis.Client

func BootRedis() *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
	})
	return client
}

func Redis() *redis.Client {
	return client
}
