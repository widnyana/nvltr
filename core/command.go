package core

import (
	"strings"

	"github.com/widnyana/nvltr/conf"
)

// Command ....
type Command struct {
	command string
	args    []string
	valid   bool
}

// ParseCommand ...
func ParseCommand(text string) Command {
	var c Command

	c.valid = strings.HasPrefix(text, conf.Config.Core.CommandPrefix)
	if c.valid {
		split := strings.Split(text, " ")
		c.command = split[0]
		c.args = split[1:]
	}

	return c
}

// Valid return true if Command contain valid command data
func (c Command) Valid() bool {
	return c.valid
}

// Cmd return valid command
func (c Command) Cmd() string {
	return c.command
}

// GetArgs return valid command arguments
func (c Command) GetArgs() string {
	return strings.Join(c.args, " ")
}

// GetArgsIdx get args until given index
func (c Command) GetArgsIdx(until int) (requested string, remainder string) {
	count := len(c.args)
	LogAccess.Infof("Length: %d, until: %d", count, until)

	if count <= until {
		requested = strings.Join(c.args, " ")
		remainder = ""
	} else {

		requested = strings.Join(c.args[0:until], " ")
		remainder = strings.Join(c.args[until+1:], " ")
	}
	return
}
