package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	/* IDENTIFIERS + LITERALS */
	INDENT = "INDENT"
	INT    = "INT"

	/* OPERATORS */
	ASSIGN = "="
	PLUS   = "+"

	/* DELIMITERS */
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
