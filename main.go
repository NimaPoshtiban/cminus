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
	var ast string
	flag.StringVar(&filename, "interpret", "", "interpret the given file")
	flag.StringVar(&ast, "ast", "", "generates program ast and save it to a file in a json format")
	flag.Parse()

	if ast != "" && filename != "" {
		repl.GeneratesAST(filename, ast)
	}

	if filename == "" {
		repl.Start(os.Stdin, os.Stdout)
		fmt.Printf("Hello %s! This is the C-- programming language!\n", user.Username)
		fmt.Println("Feel free to type in commands")

	} else {
		repl.Interpret(filename, os.Stdout)
	}
}
