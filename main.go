package main

import (
	"fmt"
	"os"

	"github.com/widnyana/nvltr/core"
)

//go:generate go run ./cert/loader.go

func main() {
	defer func() {
		if r := recover(); r != nil {
			core.LogError.Error("Recovered in f")
		}
	}()

	err := core.Boot()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(-1)
	}

	err = core.InitBot()
	if err != nil {
		core.LogError.Errorf("Error on initing BOT. %s", err.Error())
		os.Exit(1)
	}
	server := core.MakeHTTPServer()

	err = server.RunServer()
	if err != nil {
		core.LogError.Errorf("Error Running Server: \n%s", err.Error())
	}
}
