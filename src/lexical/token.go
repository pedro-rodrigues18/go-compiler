package lexical

type TokenType int

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	NUMBER
	SYMBOL
	STRING
	CHARACTER
	EOF
)

type Token struct {
	Type  TokenType
	Value string
}
