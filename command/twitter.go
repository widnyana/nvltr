package command

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/url"
	"strings"

	"github.com/redite/tlbot"
	"github.com/widnyana/nvltr/conf"
	"github.com/widnyana/nvltr/util"
)

func init() {
	register(cmdPostTweet)
	register(cmdGetTL)
	register(cmdTwtMention)
}

var (
	cmdGetTL = &Command{
		Name:      "tl",
		ShortLine: "Get User Timeline. [/tl nvltr]",
		Run:       getUserTl,
	}
	cmdPostTweet = &Command{
		Name:      "tweet",
		ShortLine: "Send tweet. [/tweet message]",
		Run:       sendTweet,
	}
	cmdTwtMention = &Command{
		Name:      "mention",
		ShortLine: "Send tweet. [/mention]",
		Run:       getTwitterNotif,
	}
	username = "your_twitter_username"
)

func getUserTl(ctx context.Context, b *tlbot.Bot, msg *tlbot.Message) {
	args := msg.Args()
	if len(args) == 0 {
		if _, err := b.SendMessage(msg.Chat.ID, "please provide a username", nil); err != nil {
			log.Printf("Error while sending message: %v\n", err)
		}
		return
	}

	username := strings.Replace(strings.Join(args, ""), " ", "", -1)
	t := conf.Config.Account.Twitter
	twttr := util.NewTwitterAPI(t.ConsumerKey, t.ConsumerSecret, t.AccessToken, t.AccessSecret)

	v := url.Values{}
	v.Add("screen_name", username)
	v.Add("count", "5")

	userTL, err := twttr.API.GetUserTimeline(v)
	if err != nil {
		b.SendMessage(msg.Chat.ID, fmt.Sprintf("dude, we got error. %s", err.Error()), nil)
		return
	}

	opts := new(tlbot.SendOptions)
	opts.ParseMode = tlbot.ModeMarkdown

	for _, t := range userTL {
		text := html.UnescapeString(t.Text)
		_, err = b.SendMessage(
			msg.Chat.ID,
			fmt.Sprintf("`[[%s]] -- @%s: %s`", t.CreatedAt[0:19], t.User.ScreenName, text),
			opts,
		)
		if err != nil {
			fmt.Printf("Error Replying: %s", err.Error())
		}
	}
}

func sendTweet(ctx context.Context, b *tlbot.Bot, msg *tlbot.Message) {
	args := msg.Args()
	if len(args) == 0 {
		if _, err := b.SendMessage(msg.Chat.ID, "please add your tweet.", nil); err != nil {
			log.Printf("Error while sending message: %v\n", err)
		}
		return
	}

	tweet := strings.Join(args, " ")
	t := conf.Config.Account.Twitter
	twttr := util.NewTwitterAPI(t.ConsumerKey, t.ConsumerSecret, t.AccessToken, t.AccessSecret)

	res, err := twttr.API.PostTweet(tweet, url.Values{})
	if err != nil {
		b.SendMessage(msg.Chat.ID, "error when submitting tweet.", nil)
		fmt.Printf("Error tweeting: %s", err.Error())
		return
	}

	b.SendMessage(msg.Chat.ID, fmt.Sprintf("https://twitter.com/%s/status/%s", username, res.IdStr), nil)
}

func getTwitterNotif(ctx context.Context, b *tlbot.Bot, msg *tlbot.Message) {

	opts := new(tlbot.SendOptions)
	opts.ParseMode = tlbot.ModeMarkdown

	t := conf.Config.Account.Twitter
	twttr := util.NewTwitterAPI(t.ConsumerKey, t.ConsumerSecret, t.AccessToken, t.AccessSecret)

	v := url.Values{}
	v.Add("count", "5")

	fmt.Println("Checking Mentions...")
	res, err := twttr.API.GetMentionsTimeline(v)
	if err != nil {
		b.SendMessage(msg.Chat.ID, "oops, error occured.", nil)
		fmt.Printf("Error GetMentionsTimeline: %s", err.Error())
	}

	for _, t := range res {
		text := html.UnescapeString(t.Text)
		_, err = b.SendMessage(
			msg.Chat.ID,
			fmt.Sprintf("`[[%s]] -- @%s: %s`", t.CreatedAt[0:19], t.User.ScreenName, text),
			opts,
		)
		if err != nil {
			fmt.Printf("Error Replying: %s", err.Error())
		}
	}
}
