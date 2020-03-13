package tgbot

import (
	"encoding/json"
	"flag"
	"github.com/go-redis/redis/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gohouse/golib/e"
	"github.com/gohouse/golib/exception"
	"github.com/gohouse/golib/tgbot/adapter"
	"github.com/gohouse/golib/tgbot/api"
	"github.com/gohouse/golib/tgbot/store"
	"log"
	"os"
	"strings"
	"sync"
)

var f string
var token string

func init() {
	flag.StringVar(&f, "f", "", "放置bot的token的配置文件")
	flag.StringVar(&token, "token", "", "bot的token,与f参数2选其一即可")
	flag.Parse()
	if token=="" {
		if f=="" {
			f="config.json"
		}
		fp,err:=os.Open(f)
		if err!=nil {
			panic(err.Error())
		}
		var tokenobj struct{
			Token string
		}
		err = json.NewDecoder(fp).Decode(&tokenobj)
		if err!=nil {
			panic(err.Error())
		}
		token = tokenobj.Token
	}
}

type Cors func(ctx *Context)

type TgBot struct {
	BotAPI   *tgbotapi.BotAPI
	updates  *tgbotapi.UpdatesChannel
	handlers []Cors
	cmds     *sync.Map
	Store    *redis.Client
	option   *Options
}

func NewTgBot(opts ...OptionHandleFunc) *TgBot {

	var opt Options
	for _, o := range opts {
		o(&opt)
	}

	// 初始化bot
	bot := getBot()

	// 初始化 updates
	updates := getUpdates(bot)

	tb := &TgBot{BotAPI: bot, updates: updates, cmds: &sync.Map{}, Store: store.BootRedis(), option: &opt}

	tb.Listen("help", api.NewHelp())

	return tb
}

func (tb *TgBot) Use(hs ...Cors) {
	tb.handlers = append(tb.handlers, hs...)
}

func (tb *TgBot) Run() {
	// 监听消息
	for update := range *tb.updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		//var cp = update
		exception.TryCatch(func() {
			go NewContext(tb, update).RunCmd()
		}, func(err e.Error) {
			log.Println(err.Error())
		})
	}
}

func (tb *TgBot) Listen(cmd string, adp adapter.ImgAdapter) *TgBot {
	cmd = strings.TrimPrefix(cmd, "/")
	tb.cmds.Store(cmd, adp)
	return tb
}

func getBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	bot.Debug = true
	if err != nil {
		panic(err.Error())
	}
	return bot
}
func getUpdates(bot *tgbotapi.BotAPI) *tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	//bot.GetUpdates(u)
	//if err != nil {
	//	log.Panic(err)
	//}
	return &updates
}
//func getWebhookUpdates(bot *tgbotapi.BotAPI) *tgbotapi.UpdatesChannel {
//	updates := bot.ListenForWebhook("/" + bot.Token)
//	return &updates
//}
