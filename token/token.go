package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT" // add , foobar ,x ,y ,...
	INT   = "INT"   // integers

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MODULO   = "%"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	STRING    = "STRING"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	LBRACKET = "["
	RBRACKET = "]"

	COLON = ":"
	// keywords
	MACRO    = "macro"
	FUNCTION = "func"
	LET      = "let"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

// language keywords
var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"macro":  MACRO,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// check to see if the given identifier is  a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
