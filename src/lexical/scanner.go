package lexical

import (
	"unicode"
)

func Scan(sourceCode string) []Token {
	var tokens []Token
	start := 0

	for i := range sourceCode {
		switch {
		case isIdentifier(sourceCode[start : i+1]):
			continue
		case unicode.IsSpace(rune(sourceCode[i])):
			if start < i {
				token := Token{Type: IDENTIFIER, Value: sourceCode[start:i]}
				tokens = append(tokens, token)
			}
			start = i + 1
		case isKeyword(sourceCode[start : i+1]):
			token := Token{Type: KEYWORD, Value: sourceCode[start : i+1]}
			tokens = append(tokens, token)
			start = i + 1

		case isString(sourceCode[start : i+1]):
			token := Token{Type: STRING, Value: sourceCode[start : i+1]}
			tokens = append(tokens, token)
		case isNumber(sourceCode[start : i+1]):
			token := Token{Type: NUMBER, Value: sourceCode[start : i+1]}
			tokens = append(tokens, token)
			start = i + 1

		default:
			token := Token{Type: SYMBOL, Value: string(sourceCode[i])}
			tokens = append(tokens, token)
			start = i + 1
		}
	}
	if start < len(sourceCode) {
		token := Token{Type: IDENTIFIER, Value: sourceCode[start:]}
		tokens = append(tokens, token)
	}

	tokens = append(tokens, Token{Type: EOF, Value: ""})

	return tokens
}

func isKeyword(value string) bool {
	keywords := map[string]bool{
		"auto":     true,
		"break":    true,
		"case":     true,
		"char":     true,
		"const":    true,
		"continue": true,
		"default":  true,
		"do":       true,
		"double":   true,
		"else":     true,
		"float":    true,
		"for":      true,
		"if":       true,
		"int":      true,
		"long":     true,
		"return":   true,
		"sizeof":   true,
		"static":   true,
		"struct":   true,
		"switch":   true,
		"typedef":  true,
		"unsigned": true,
		"void":     true,
		"while":    true,
	}

	return keywords[value]
}

func isIdentifier(value string) bool {
	if !unicode.IsLetter(rune(value[0])) && value[0] != '_' {
		return false
	}
	for _, ch := range value {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' {
			return false
		}
	}
	return true
}

func isNumber(value string) bool {
	for _, r := range value {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func isString(value string) bool {
	if value[0] != '"' || value[len(value)-1] != '"' {
		return false
	}
	return true
}
