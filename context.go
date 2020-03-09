package tgbot

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gohouse/tgbot/adapter"
	"log"
	"math"
	"strings"
	"time"
)

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	TgBot    *TgBot
	Update   tgbotapi.Update
	index    int8
	allAlive int
}
type ContextHandleFunc func(ctx *Context)

func SetAllAlive(alive int) ContextHandleFunc {
	return func(ctx *Context) {
		ctx.allAlive = alive
	}
}

func NewContext(tb *TgBot, update tgbotapi.Update) *Context {
	return &Context{TgBot: tb, Update: update}
}

func (c *Context) RunCmd() {
	// 先走中间件
	c.start()

	// 再走命令
	if c.IsAborted() {
		return
	}

	// 获取命令
	if c.Update.Message.IsCommand() {
		var cfg tgbotapi.Chattable
		cmd := c.Update.Message.Command()
		cmdArgs := c.Update.Message.CommandArguments()
		if r, ok := c.TgBot.cmds.Load(cmd); ok {
			adp := r.(adapter.ImgAdapter)
			// 如果是 /help, 特殊处理
			switch cmd {
			case "help":
				cfg = c.buildHelp(c.Update.Message.Chat.ID)
			default:
				cfg = adp.BuildChattable(&c.Update, cmdArgs)
			}
			var msg tgbotapi.Message
			var err error
			if cmd == "mm" { // 远程
				msg, err = send(c.TgBot.BotAPI, cfg)
			} else { // 其他
				msg, err = c.TgBot.BotAPI.Send(cfg)
			}
			//if cmd != "help" &&
			//	cmd != "neihan" &&
			//	cmd != "gif" &&
			//	cmd != "kfc" &&
			//	cmd != "asian" &&
			//	cmd != "cc" &&
			//	cmd != "duanzi" { // 远程
			//	msg, err = send(c.TgBot.BotAPI, cfg)
			//} else { // 本地
			//	msg, err = c.TgBot.BotAPI.Send(cfg)
			//}
			if err != nil {
				log.Println("send err:", err.Error())
			} else {
				// 撤回消息
				go destroyMsg(c, msg, c.Update, cmd, adp.Alive())
			}
		} else {
			//js,_:=json.Marshal(c.Update.Message)
			//fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
			////fmt.Println(string(js))
			//file.FilePutContents("test.json",js,os.O_TRUNC)
			//fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
			//c.TgBot.BotAPI.Send(tgbotapi.NewDeleteMessage(c.Update.Message.Chat.ID,c.Update.Message.ReplyToMessage.MessageID))

			cfg = tgbotapi.NewMessage(c.Update.Message.Chat.ID, "不支持的命令: /"+cmd)
			_, err := c.TgBot.BotAPI.Send(cfg)
			if err != nil {
				log.Println("send err:", err.Error())
			}
		}

	}
}

//destroyMsg 销毁图片
func destroyMsg(c *Context, msg tgbotapi.Message, update tgbotapi.Update, cmd string, alive int64) {
	if c.TgBot.option.alive > 0 {
		alive = c.TgBot.option.alive
	}
	if alive == 0 || cmd == "help" {
		return
	}
	time.Sleep(time.Duration(alive) * time.Second)
	//tb.bot
	c.TgBot.BotAPI.Send(tgbotapi.NewDeleteMessage(msg.Chat.ID, msg.MessageID))
	msg2, _ := c.TgBot.BotAPI.Send(tgbotapi.NewMessage(msg.Chat.ID,
		fmt.Sprintf("超时销毁 %s 发送的命令消息: /%s", update.Message.From.FirstName, cmd)))
	//tb.bot.DeleteChatPhoto(tgbotapi.DeleteChatPhotoConfig{ChatID: msg.Chat.ID})
	go func() {
		time.Sleep(5 * time.Second)
		if msg2.Chat.ID != 0 && msg2.MessageID != 0 {
			c.TgBot.BotAPI.Send(tgbotapi.NewDeleteMessage(msg2.Chat.ID, msg2.MessageID))
		}
		if update.Message.Chat.ID != 0 && update.Message.MessageID != 0 {
			// 撤回发信人命令消息
			c.TgBot.BotAPI.Send(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
		}
	}()
}

func send(bot *tgbotapi.BotAPI, c tgbotapi.Chattable) (tgbotapi.Message, error) {
	resp, err := bot.Request(c)
	if err != nil {
		return tgbotapi.Message{}, err
	}

	var message []tgbotapi.Message
	err = json.Unmarshal(resp.Result, &message)

	if err != nil {
		return tgbotapi.Message{}, err
	} else {
		if len(message) > 0 {
			return message[0], nil
		}
	}
	return tgbotapi.Message{}, err
}

/************************************/
/*********** FLOW CONTROL ***********/
/************************************/

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in GitHub.
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.TgBot.handlers)) {
		c.TgBot.handlers[c.index](c)
		c.index++
	}
}
func (c *Context) start() {
	c.index = 0
	if len(c.TgBot.handlers) > 0 {
		c.TgBot.handlers[c.index](c)
		c.Next()
	}
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) buildHelp(chatid int64) tgbotapi.Chattable {
	var cmdlist = []string{fmt.Sprintf("已支持的命令如下, 为了界面清洁, 图片会在60s后自动撤回, 建议将我设置为管理员以便提供更好的服务:")}
	c.TgBot.cmds.Range(func(k, v interface{}) bool {
		adp := v.(adapter.ImgAdapter)
		cmdlist = append(cmdlist, fmt.Sprintf("/%s - %s", k, adp.Description()))
		return true
	})
	return tgbotapi.NewMessage(chatid, strings.Join(cmdlist, "\n"))
}
