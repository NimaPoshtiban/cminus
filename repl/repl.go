package repl

import (
	"bufio"
	"fmt"
	"interpreter/evaluator"
	"interpreter/lexer"
	"interpreter/object"
	"interpreter/parser"
	"io"
)

const C_MINUS_MINUS = `                                                                                                       
 _____ _     _  ____
/ ___/| |   | |/ ___| 
| (___| |__ | |\___ \ 
\___ \| '_ \| | ___) |
 ___) | | | | |/ ___/ 
|____/|_| |_|_/_/                              
											   `
const Prompt = "-> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	lineNumber := 0
	for {
		lineNumber++
		fmt.Printf("%d) %s ", lineNumber, Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, C_MINUS_MINUS)
	io.WriteString(out, "\nWoops! We ran into some C-- business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\t\n")
	}
}
