package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the C-- programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	// Check if there is any input from standard input.
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == 0 {
		// There is input from a pipe or a file.
		file := os.Stdin

		// Pass the file object to the REPL.
		repl.Interpret(file, os.Stdout)
	} else {
		// There is no input from a pipe or a file.
		repl.Start(os.Stdin, os.Stdout)
	}
}
