package repl

import (
	"bufio"
	"cminus/evaluator"
	"cminus/lexer"
	"cminus/object"
	"cminus/parser"
	"fmt"
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

// Start reads input from the given io.Reader and writes output to the given io.Writer.
//
// Parameters:
// - in: the input reader to read from.
// - out: the output writer to write to.
func Start(in io.Reader, out io.Writer) {

	green := "\033[32m"
	reset := "\033[0m"

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		fmt.Fprintf(out, "%d) %s ", lineNumber, Prompt)
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		evaluator.DefineMacros(program, macroEnv)

		if len(p.Errors()) != 0 {
			errs := p.Errors()
			errs = append(errs, fmt.Sprintf("line %d", lineNumber))
			printParseErrors(out, errs)
			continue
		}
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)

		evaluated = evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Fprintf(out, green+"%s"+reset+"\n", evaluated.Inspect())
		}
	}
}

// Interpret reads input from a reader and evaluates it line by line.
//
// It takes two parameters:
// - input: an io.Reader from which the input is read.
// - output: an io.Writer to which the output is written.
//
// The function does the following:
// - Creates a scanner to read input line by line.
// - Creates a new environment to store variables and functions.
// - Loops through each line of input.
//   - Creates a lexer for the line.
//   - Creates a parser for the lexer.
//   - Parses the program and returns the AST.
//   - If there are any parsing errors, writes the first error to the output.
//   - Evaluates the AST using the evaluator.
//
// The function does not return anything.
func Interpret(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	env := object.NewEnvironment()

	for scanner.Scan() {
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			fmt.Fprintf(output, "Error: %s\n", p.Errors()[0])
			continue
		}

		_ = evaluator.Eval(program, env)
	}
}

func printParseErrors(out io.Writer, errors []string) {
	red := "\033[31m"
	reset := "\033[0m"

	io.WriteString(out, red+C_MINUS_MINUS+reset)
	io.WriteString(out, "\nWoops! We ran into some C-- business here!\n")
	io.WriteString(out, red+" parser errors:"+reset+"\n")
	for _, msg := range errors {
		io.WriteString(out, red+"\t"+msg+"\t"+reset+"\n")
	}
}
