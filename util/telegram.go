package util

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"github.com/k0kubun/pp"
	"gopkg.in/telegram-bot-api.v4"
)

// TBot hold Telegram Bot data
type TBot struct {
	Bot     *tgbotapi.BotAPI
	Updates chan tgbotapi.Update
	handler gin.HandlerFunc
}

// TGBot the global telegram bot
var TGBot *TBot

// NewTGBot build Telegram Bot
func NewTGBot(token, webhookPath, webhookURL, cert string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		pp.Print(err.Error())
	}

	bot.Debug = false
	_, err = bot.SetWebhook(
		tgbotapi.NewWebhookWithCert(webhookURL, cert),
	)
	if err != nil {
		panic(err.Error())
	}

	TGBot.Bot = bot
	TGBot.Updates = make(chan tgbotapi.Update, 100)
	TGBot.handler = func(c *gin.Context) {
		bytes, _ := ioutil.ReadAll(c.Request.Body)

		var update tgbotapi.Update
		json.Unmarshal(bytes, &update)

		TGBot.Updates <- update
	}
}

// GetHTTPHandler return HTTP handler for webhook
func (t *TBot) GetHTTPHandler() gin.HandlerFunc {
	return t.handler
}
