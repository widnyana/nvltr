package command

import (
	"context"
	"strings"

	"fmt"

	"github.com/redite/tlbot"
	"github.com/widnyana/nvltr/util"
)

func init() {
	register(cmdWhois)
}

var cmdWhois = &Command{
	Name:      "ws",
	ShortLine: "check whois [/ws domname.tld]",
	Run:       runGetWhois,
}

func runGetWhois(ctx context.Context, b *tlbot.Bot, msg *tlbot.Message) {
	opts := new(tlbot.SendOptions)
	opts.ParseMode = tlbot.ModeNone

	args := msg.Args()

	if len(args) == 0 {
		b.SendMessage(msg.Chat.ID, "format: `/ws domname.tld`.", opts)
		return
	}

	domname := strings.Join(args, "")
	w, err := util.Whois(domname)
	if err != nil {
		b.SendMessage(msg.Chat.ID, "sorry.", opts)
		fmt.Printf("dude, we got error: %s", err.Error())
		return
	}

	info, err := util.WhoisParser(w)
	if err != nil {
		b.SendMessage(msg.Chat.ID, "couldn't decode response.", opts)
		return
	}

	reply := fmt.Sprintf(`
		Domain Info
===========
Domain: %s
Creation: %s
Expires: %s
NameServers: %s
Registrar: %s
Email: %s
	`,
		info.Registrar.DomainName,
		info.Registrar.CreatedDate,
		info.Registrar.ExpirationDate,
		info.Registrar.NameServers,
		info.Registrant.Name,
		info.Registrant.Email,
	)
	reply = fmt.Sprintf("`%s`", reply)

	opts.ParseMode = tlbot.ModeMarkdown
	b.SendMessage(msg.Chat.ID, reply, opts)
}
