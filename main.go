package main

import (
	"cminus/repl"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s! This is the C-- programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands")

	if !isInputFromPipe() {
		repl.Start(os.Stdin, os.Stdout)
	} else {
		file := os.Stdin
		repl.Interpret(file, os.Stdout)
	}
}

func isInputFromPipe() bool {
	info, _ := os.Stdin.Stat()
	return (info.Mode() & os.ModeCharDevice) == 0
}
