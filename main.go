package main

import (
	"cminus/repl"
	"flag"
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
	var filename string
	flag.StringVar(&filename, "interpret", "", "interpret the given file")
	flag.Parse()

	if filename == "" {
		repl.Start(os.Stdin, os.Stdout)
		fmt.Printf("Hello %s! This is the C-- programming language!\n", user.Username)
		fmt.Println("Feel free to type in commands")

	} else {
		repl.Interpret(filename, os.Stdout)
	}
}

func isInputFromPipe() bool {
	info, _ := os.Stdin.Stat()
	return (info.Mode() & os.ModeCharDevice) == 0
}
