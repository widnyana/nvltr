package command

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/redite/tlbot"
)

func init() {
	register(cmdNsLookup)
}

var cmdNsLookup = &Command{
	Name:      "nslookup",
	ShortLine: "Name Server Lookup [/nslookup domname.tld]",
	Run:       mynslookup,
}

func mynslookup(ctx context.Context, b *tlbot.Bot, msg *tlbot.Message) {
	opts := new(tlbot.SendOptions)
	opts.ReplyTo = msg.ID
	opts.ParseMode = tlbot.ModeMarkdown

	args := msg.Args()
	if len(args) == 0 {
		if _, err := b.SendMessage(msg.Chat.ID, "nyari apa bos?", nil); err != nil {
			log.Printf("Error while sending message: %v\n", err)
		}
		return
	}

	dom := strings.Join(args, "")
	hs, err := net.LookupHost(dom)
	if err != nil {
		b.SendMessage(msg.Chat.ID, err.Error(), opts)
		return
	}

	reply := strings.Join(hs, "\n")
	b.SendMessage(msg.Chat.ID, reply, opts)
}
