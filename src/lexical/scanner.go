package lexical

import (
	"unicode"
)

func Scan(sourceCode string) []Token {
	var tokens []Token
	start := 0

	for i := 0; i < len(sourceCode); i++ {
		for start < len(sourceCode) && unicode.IsSpace(rune(sourceCode[start])) {
			start++
		}
		i = start

		if i >= len(sourceCode) {
			break
		}

		switch {
		case unicode.IsLetter(rune(sourceCode[i])) || sourceCode[i] == '_':
			for i < len(sourceCode) && (unicode.IsLetter(rune(sourceCode[i])) || unicode.IsDigit(rune(sourceCode[i])) || sourceCode[i] == '_') {
				i++
			}
			value := sourceCode[start:i]
			if isKeyword(value) {
				tokens = append(tokens, Token{Type: KEYWORD, Value: value})
			} else {
				tokens = append(tokens, Token{Type: ID, Value: value})
			}
			start = i
		case unicode.IsDigit(rune(sourceCode[i])):
			for i < len(sourceCode) && unicode.IsDigit(rune(sourceCode[i])) {
				i++
			}
			tokens = append(tokens, Token{Type: NUMBER, Value: sourceCode[start:i]})
			start = i
		case sourceCode[i] == '"':
			j := i + 1
			for j < len(sourceCode) && sourceCode[j] != '"' {
				j++
			}
			if j < len(sourceCode) {
				tokens = append(tokens, Token{Type: STRING, Value: sourceCode[start : j+1]})
				i = j
			}
			start = i + 1
		case isSymbol(string(sourceCode[i])):
			tokens = append(tokens, Token{Type: SYMBOL, Value: string(sourceCode[i])})
			start = i + 1
		case unicode.IsSpace(rune(sourceCode[i])):
			start = i + 1
		}
	}

	if start < len(sourceCode) {
		tokens = append(tokens, Token{Type: ID, Value: sourceCode[start:]})
	}

	tokens = append(tokens, Token{Type: EOF, Value: "$"})

	return tokens
}

func isSymbol(value string) bool {
	symbols := map[string]bool{
		"+": true, "-": true, "*": true, "/": true,
		"=": true, "<": true, ">": true, "(": true, ")": true,
		"{": true, "}": true, "[": true, "]": true, ";": true, ",": true,
	}
	return symbols[value]
}

func isKeyword(value string) bool {
	keywords := map[string]bool{
		"int":     true,
		"float":   true,
		"char":    true,
		"boolean": true,
		"void":    true,
		"if":      true,
		"else":    true,
		"for":     true,
		"while":   true,
		"scanf":   true,
		"printf":  true,
		"main":    true,
		"return":  true,
	}

	return keywords[value]
}
