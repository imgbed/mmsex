package tgbot
//
//import (
//	"encoding/json"
//	"flag"
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
//	"github.com/gohouse/golib/e"
//	"github.com/gohouse/golib/exception"
//	"github.com/gohouse/tgbot/adapter"
//	"github.com/gohouse/tgbot/api"
//	"log"
//	"net/http"
//	"os"
//	"strings"
//	"sync"
//)
//
//var f string
//var token string
//
//func init() {
//	flag.StringVar(&f, "f", "", "放置bot的token的配置文件")
//	flag.StringVar(&token, "token", "", "bot的token,与f参数2选其一即可")
//	flag.Parse()
//	if token=="" {
//		if f=="" {
//			f="config.json.example"
//		}
//		fp,err:=os.Open(f)
//		if err!=nil {
//			panic(err.Error())
//		}
//		var tokenobj struct{
//			Token string
//		}
//		err = json.NewDecoder(fp).Decode(&tokenobj)
//		if err!=nil {
//			panic(err.Error())
//		}
//		token = tokenobj.Token
//	}
//}
//
//type Cors func(ctx *Context)
//
//type TgBot struct {
//	BotAPI   *tgbotapi.BotAPI
//	updates  *tgbotapi.UpdatesChannel
//	handlers []Cors
//	cmds     *sync.Map
//	option   *Options
//}
//
//func NewTgBot()  {
//	bot, err := tgbotapi.NewBotAPI(token)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	updates := bot.ListenForWebhook("/" + bot.Token)
//	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
//
//	for update := range updates {
//		log.Printf("%+v\n", update)
//	}
//}
//
//func NewTgBot2(opts ...OptionHandleFunc) *TgBot {
//	var opt Options
//	for _, o := range opts {
//		o(&opt)
//	}
//
//	// 初始化bot
//	bot := getBot()
//
//	// 初始化 updates
//	updates := getUpdates(bot)
//
//	tb := &TgBot{BotAPI: bot, updates: updates, cmds: &sync.Map{}, option: &opt}
//
//	tb.Listen("help", api.NewHelp())
//
//	return tb
//}
//
//func (tb *TgBot) Use(hs ...Cors) {
//	tb.handlers = append(tb.handlers, hs...)
//}
//
//func (tb *TgBot) Run() {
//	// 监听消息
//	for update := range *tb.updates {
//		if update.Message == nil { // ignore any non-Message Updates
//			continue
//		}
//
//		//var cp = update
//		exception.TryCatch(func() {
//			go NewContext(tb, update).RunCmd()
//		}, func(err e.Error) {
//			log.Println(err.Error())
//		})
//	}
//}
//
//func (tb *TgBot) Listen(cmd string, adp adapter.ImgAdapter) *TgBot {
//	cmd = strings.TrimPrefix(cmd, "/")
//	tb.cmds.Store(cmd, adp)
//	return tb
//}
//
//func getBot() *tgbotapi.BotAPI {
//	bot, err := tgbotapi.NewBotAPI(token)
//	bot.Debug = true
//	if err != nil {
//		panic(err.Error())
//	}
//	return bot
//}
//func getUpdates(bot *tgbotapi.BotAPI) *tgbotapi.UpdatesChannel {
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//	updates := bot.GetUpdatesChan(u)
//	//bot.GetUpdates(u)
//	//if err != nil {
//	//	log.Panic(err)
//	//}
//	return &updates
//}
