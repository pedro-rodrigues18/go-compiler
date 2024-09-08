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
