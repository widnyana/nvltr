package core

import (
	"github.com/gin-gonic/gin"
	"github.com/widnyana/nvltr/conf"
)

type route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

func routeProvider() []route {
	return []route{
		{
			method:  "POST",
			path:    conf.Config.Account.Telegram.Webhook,
			handler: WebhookHandler,
		},
	}
}
