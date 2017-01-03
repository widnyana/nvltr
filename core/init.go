package core

import "github.com/widnyana/nvltr/conf"

// InitBot functionality
func InitBot() error {
	_, err := NewTGBot(conf.Config.Account.Telegram)
	if err != nil {
		LogError.Errorf("Building TGBot: failed. %s", err.Error())
	}

	return nil
}
