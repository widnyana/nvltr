package core

import "testing"

func TestParseCommand(t *testing.T) {
	InitLog()

	scenario := []struct {
		text    string
		command string
		args    string
		withIdx string
	}{
		{
			text:    "!tweet haloooooo",
			command: "!tweet",
			args:    "haloooooo",
			withIdx: "haloooooo",
		},
		{
			text:    "!tweet 1234 lalalala",
			command: "!tweet",
			args:    "1234 lalalala",
			withIdx: "1234",
		},
	}

	for _, idx := range scenario {
		c := ParseCommand(idx.text)
		if c.Cmd() != idx.command {
			t.Errorf("Command doesnt match. expecting: '%s' got: '%s'", idx.command, c.Cmd())
		}

		if c.GetArgs() != idx.args {
			t.Errorf("Args doesnt match. expecting: '%s' got: '%s'", idx.args, c.GetArgs())
		}

		requested, _ := c.GetArgsIdx(1)
		if requested != idx.withIdx {
			t.Errorf("ArgsWithIdx doesnt match. expecting: '%s' got: '%s'", idx.withIdx, requested)
		} else {
			t.Logf("Requested: %s, got: %s", idx.withIdx, requested)
		}
	}
}
