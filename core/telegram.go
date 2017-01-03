package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/redite/tlbot"
	"github.com/widnyana/nvltr/command"
	"github.com/widnyana/nvltr/conf"
)

var (
	bot       tlbot.Bot
	messageCh chan tlbot.Message
)

func init() {
	messageCh = make(chan tlbot.Message, 100)
}

// NewTGBot create new tg bot
func NewTGBot(conf conf.Telegram) (tlbot.Bot, error) {
	bot = tlbot.New(conf.Token)

	if err := bot.SetWebhook(fmt.Sprintf("%s/%s", conf.BaseURL, conf.Webhook)); err != nil {
		return bot, fmt.Errorf("Cannot Set Webhook: %s", err.Error())
	}

	fmt.Printf("Success registering Webhook at: %s\n", fmt.Sprintf("%s/%s", conf.BaseURL, conf.Webhook))
	return bot, nil
}

func handleIncoming(ctx context.Context, ch <-chan tlbot.Message) {
	fmt.Print("[handleIncoming] Listening on incoming request.\n")
	for msg := range ch {

		fmt.Printf("found msg: %s from: %d\n", msg.Text, msg.Chat.ID)

		// react only to user sent messages
		if msg.IsService() {
			fmt.Print("msg is not service")
			continue
		}
		// is message a bot command?
		cmdname := msg.Command()
		if cmdname == "" {
			fmt.Print("msg is not command")
			continue
		}

		// is the command even registered?
		cmd := command.Lookup(cmdname)
		if cmd == nil {
			continue
		}

		// it is. cool, run it!
		go cmd.Run(ctx, &bot, &msg)
	}
}

func newCtxWithValues(c conf.Conf) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, conf.CtxGoogleAPIKey, c.Account.GoogleMapsKey)
	ctx = context.WithValue(ctx, conf.CtxTwitterConf, c.Account.Twitter)
	return ctx
}

// WebhookHandler ...
func WebhookHandler(ctx *gin.Context) {
	runtime.Gosched()
	go handleIncoming(cfx, messageCh)

	var u tlbot.Update
	if err := json.NewDecoder(ctx.Request.Body).Decode(&u); err != nil {
		fmt.Printf("error decoding request body: %v\n", err)
		ctx.String(http.StatusBadRequest, "failed")
		ctx.Done()
	}
	ctx.String(http.StatusOK, "done")

	messageCh <- u.Payload
}
