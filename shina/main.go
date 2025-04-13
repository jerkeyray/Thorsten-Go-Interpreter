package main

import (
	"fmt"
	"os"
	"os/user"
	"shina/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Shina Programming Language!\n", user.Username)
	fmt.Printf("Feel free to type in commands!\n")
	repl.Start(os.Stdin, os.Stdout)

}
