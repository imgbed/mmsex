package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	//port      = os.Getenv("PORT")
	publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
	token     = os.Getenv("TOKEN")      // you must add it to your config vars
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert(fmt.Sprintf("%s:443/%s", publicURL, token), "cert/cert.pem"))
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(fmt.Sprintf("%s/%s", publicURL, token)))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + token)
	for update := range updates {
		log.Printf("%+v\n", update)
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.From.FirstName))
	}
}
