package util

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// TwitterAPI hold twitter API struct
type TwitterAPI struct {
	v   url.Values
	API *anaconda.TwitterApi
}

// Twitter hold global TwitterAPI instance
var Twitter *TwitterAPI

// NewTwitterAPI provide mechanism to build a twitterAPI instance
func NewTwitterAPI(ck, cs, at, as string) *TwitterAPI {
	Twitter = new(TwitterAPI)

	anaconda.SetConsumerKey(ck)
	anaconda.SetConsumerSecret(cs)

	Twitter.v = url.Values{}
	Twitter.API = anaconda.NewTwitterApi(
		at,
		as,
	)

	return Twitter
}

// SetStatus Post a status
func (t *TwitterAPI) SetStatus(status string) (anaconda.Tweet, error) {
	tweet, err := t.API.PostTweet(status, t.v)

	return tweet, err
}

// Reply provide ability to reply to status
func (t *TwitterAPI) Reply(status, statusID string) (anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("in_reply_to_status_id", statusID)
	tweet, err := t.API.PostTweet(status, v)

	return tweet, err
}
