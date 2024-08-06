package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/brythnl/alpha/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s, Welcome to Alpha\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
