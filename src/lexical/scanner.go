package lexical

import (
	"unicode"
)

func Scan(sourceCode string) []Token {
	var tokens []Token
	start := 0
	length := len(sourceCode)

	for start < length {
		char := rune(sourceCode[start])

		if unicode.IsSpace(char) {
			start++
			continue
		}

		if isIdentifierChar(char) {
			end := start + 1
			for end < length && isIdentifierChar(rune(sourceCode[end])) {
				end++
			}
			value := sourceCode[start:end]
			tokenType := IDENTIFIER
			if isKeyword(value) {
				tokenType = KEYWORD
			}
			tokens = append(tokens, Token{Type: tokenType, Value: value})
			start = end
			continue
		}

		if unicode.IsDigit(char) {
			end := start + 1
			dotCount := 0
			for end < length && (unicode.IsDigit(rune(sourceCode[end])) || (rune(sourceCode[end]) == '.' && dotCount == 0)) {
				if rune(sourceCode[end]) == '.' {
					dotCount++
				}
				end++
			}
			value := sourceCode[start:end]
			tokenType := INTEGER
			if dotCount == 1 {
				tokenType = DECIMAL
			}
			tokens = append(tokens, Token{Type: tokenType, Value: value})
			start = end
			continue
		}

		if char == '"' {
			end := start + 1
			for end < length && rune(sourceCode[end]) != '"' {
				end++
			}
			if end < length {
				end++
			}
			value := sourceCode[start:end]
			tokens = append(tokens, Token{Type: STRING, Value: value})
			start = end
			continue
		}

		tokens = append(tokens, Token{Type: SYMBOL, Value: string(char)})
		start++
	}

	tokens = append(tokens, Token{Type: EOF, Value: ""})
	return tokens
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

func isIdentifierChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_'
}
