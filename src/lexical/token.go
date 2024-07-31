package lexical

type TokenType int

const (
	KEYWORD    TokenType = iota // 0
	IDENTIFIER                  // 1
	INTEGER                     // 2
	DECIMAL                     // 3
	SYMBOL                      // 4
	STRING                      // 5
	CHARACTER                   // 6
	EOF                         // 7
)

type Token struct {
	Type  TokenType
	Value string
}
