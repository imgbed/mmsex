package apiv2

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/random"
	"io/ioutil"
	"log"
	"net/http"
)


type SFZ struct {
	cooldown int
}

func NewSFZ(cooldownduration ...int) *SFZ {
	var cooldown int
	if len(cooldownduration)>0 {
		cooldown = cooldownduration[0]
	}
	return &SFZ{cooldown:cooldown}
}

func (obj *SFZ) Alive() int64 {
	return 60
}

func (mm *SFZ) Description() string {
	return "死肥宅"
}

func (mm *SFZ) GetCooldown() int {
	return mm.cooldown
}

func (mm *SFZ) BuildChattable(update *tgbotapi.Update,cmdArgs string) (tc tgbotapi.Chattable) {
	urls := []sfzHandleFunc{
		getSFZ,
		//getSFZ2,
		getSFZ3,
		getSFZ4,
	}
	idx := random.RandBetween(0,len(urls)-1)
	log.Println("idx000000000000000: - ",idx)
	img := (urls[idx])()
	tc = tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{
		tgbotapi.NewInputMediaPhoto(img),
	})
	return
}

type sfzHandleFunc func() string
func getSFZ() (str string) {
	// 获取 https://konachan.com/data/preview/b5/05/b5052f72151d565cf4a75c3ad165a98b.jpg
	var page = random.RandBetween(1,14333)
	resp, err := http.Get(fmt.Sprintf("https://konachan.com/post.json?tags=ass&limit=1&page=%v",page))
	if err != nil {
		log.Println(err.Error())
		return
	}
	b,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var res = []map[string]interface{}{}
	json.Unmarshal(b, &res)
	if r,ok:=res[0]["jpeg_url"];ok{
		return r.(string)
	}
	return
}
func getSFZ2() (str string) {
	// 获取 https://img.xjh.me/desktop/img/57441136_p0.jpg
	resp, err := http.Get("https://img.xjh.me/random_img.php?return=json")
	if err != nil {
		log.Println(err.Error())
		return
	}
	b,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var res = map[string]interface{}{}
	json.Unmarshal(b, &res)
	if r,ok:=res["img"];ok{
		return "http://"+r.(string)
	}
	return
}
func getSFZ3() (str string) {
	// 获取 https://tva3.sinaimg.cn/large/0072Vf1pgy1foxkfob8yfj31hc0u0ts9.jpg
	resp, err := http.Get("http://www.dmoe.cc/random.php?return=json")
	if err != nil {
		log.Println(err.Error())
		return
	}
	b,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var res = map[string]interface{}{}
	json.Unmarshal(b, &res)
	if r,ok:=res["imgurl"];ok{
		return r.(string)
	}
	return
}
func getSFZ4() (str string) {
	// 获取 https://konachan.net/data/preview/e9/21/e921fac5816e142f0ddfdf94900217eb.jpg
	resp, err := http.Get("https://img.paulzzh.tech/touhou/random?type=json")
	if err != nil {
		log.Println(err.Error())
		return
	}
	b,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var res = map[string]interface{}{}
	json.Unmarshal(b, &res)
	if r,ok:=res["url"];ok{
		return r.(string)
	}
	return
}
