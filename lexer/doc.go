/*
	Lexer Documentation

The package lexer has a single exported struct type called Lexer. It represents the lexical analyzer, which takes source code as input and identifies a series of tokens, which are passed on to the parser for building an Abstract Syntax Tree (AST).

# Data Structure

	type Lexer struct {
	    input        string // The input source code
	    position     int    // The current position in input (points to the current char)
	    readPosition int    // The current reading position in input (after the current char)
	    ch           byte   // The current char under examination
	}

Lexer is composed of four fields:

input: This is the input source code.
position: Current position indicates the next character that hasn't been analyzed yet.
readPosition: It points to the next character that Lexer will read.
ch: The current character that is examined.
Methods
New(input string) *Lexer
This is a Constructor method that returns a pointer to a new Lexer struct, initialized with the input source code.

NextToken() token.Token
This method is responsible for identifying the next non-whitespace token by scanning the input string. It switches on the current character l.ch and generates a corresponding token. If the character is not recognized, NextToken() creates an ILLEGAL token.

It returns a Token struct that contains the Type and Literal. The Type field indicates the type of the token, such as identifier token LET, operator token PLUS, keyword token IF, etc. The Literal field indicates the literal value of the token, like the name of an identifier, the value of an integer, or the contents of a string.

readChar()
This private method reads the next character from the input string into the ch field in the lexer. The method then updates the current position to the current reading position and increments the current reading position.

peekChar() byte
This private method looks ahead one character and returns it without consuming it.

skipWhitespace()
This method consumes all whitespaces such as spaces, tabs, and newlines between characters in the input string.

readIdentifier() string
This private method reads identifiers until it encounters a non-letter character. It gets the starting position of the identifier, reads everything to the right of it, and returns the entire identifier string.

readNumber() string
This private method reads the numerical values in the input string. It reads everything up to the delimiter character (anything that's not a number) and returns the numeric string.

readString() string
This private method reads a string enclosed in double quotation marks, " ", until it hits another double-quote or the end of the file.

newToken(tokenType token.TokenType, ch byte) token.Token
This function returns a new Token struct, given the Token type and Literal value.

isLetter(ch byte) bool
This private helper function checks if a byte is a valid letter used in programming languages—valid letters include uppercase or lowercase English letters, the underscore character, or the dollar sign.

isDigit(ch byte) bool
This private helper function checks if a byte is a decimal digit (0-9).
*/
package lexer
