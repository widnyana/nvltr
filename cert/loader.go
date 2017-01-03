package main

import (
	"fmt"
	"io"
	"os"
)

// create new directory to store your TLSKey and TLSCert
// put your fullchain.pem and privkey.pem there.
//
// then change value of `domName` to your newly-created
// directory name.
var (
	domName = "your.domain.name"
	content = fmt.Sprintf(
		"package core\n\n// domain: %s\n\n\nconst (\n", domName,
	)
)

func main() {

	out, err := os.Create("core/cert.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	out.Write([]byte(content))
	out.Write([]byte("TLSCert = `"))
	thefile, err := os.Open(fmt.Sprintf("./cert/%s/fullchain.pem", domName))
	if err != nil {
		panic(err)
	}
	io.Copy(out, thefile)
	thefile.Close()
	out.Write([]byte("`\n\n"))

	out.Write([]byte("TLSKey = `"))
	thefile, err = os.Open(fmt.Sprintf("./cert/%s/privkey.pem", domName))
	if err != nil {
		panic(err)
	}
	io.Copy(out, thefile)
	out.Write([]byte("`\n)\n"))
}
