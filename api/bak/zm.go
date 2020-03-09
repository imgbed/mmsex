package bak

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
)

type ZM struct {
	cooldown int
}

func NewZM(cooldownduration ...int) *ZM {
	var cooldown int
	if len(cooldownduration)>0 {
		cooldown = cooldownduration[0]
	}
	return &ZM{cooldown:cooldown}
}

func (mm *ZM) Description() string {
	return "正经妹妹图"
}

func (mm *ZM) GetCooldown() int {
	return mm.cooldown
}

func (mm *ZM) BuildChattable(update *tgbotapi.Update,cmdArgs string) (tc tgbotapi.Chattable) {
	img := getZMUrl()
	tc = tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{
		tgbotapi.NewInputMediaPhoto(img),
	})
	return
}


func getZMUrl() (str string) {
	// 获取
	resp, err := http.Get("https://api.ooopn.com/image/beauty/api.php?type=json")
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