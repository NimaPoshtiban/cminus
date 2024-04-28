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
	LSHIFT   = "<<"
	RSHIFT   = ">>"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	LOGICAL_AND = "and"
	LOGICAL_OR  = "or"
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
	"and":    LOGICAL_AND,
	"or":     LOGICAL_OR,
}

// check to see if the given identifier is  a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
