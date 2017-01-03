package command

import (
	"net/http"
	"time"

	"github.com/redite/tlbot"
)

var (
	httpclient = &http.Client{Timeout: 10 * time.Second}
	replyOpts  = new(tlbot.SendOptions)
)
