package syntax

import (
	"fmt"
	"strings"

	"github.com/pedro-rodrigues18/go-compiler/src/lexical"
)

var parsingTable = map[string]map[string]string{
	"PROGRAMA": {
		"int":     "LISTAFUNÇÕES PRINCIPAL",
		"float":   "LISTAFUNÇÕES PRINCIPAL",
		"char":    "LISTAFUNÇÕES PRINCIPAL",
		"boolean": "LISTAFUNÇÕES PRINCIPAL",
		"void":    "LISTAFUNÇÕES PRINCIPAL",
		"main":    "LISTAFUNÇÕES PRINCIPAL",
	},
	"LISTAFUNÇÕES": {
		"int":     "DECFUNÇÃO LISTAFUNÇÕES",
		"float":   "DECFUNÇÃO LISTAFUNÇÕES",
		"char":    "DECFUNÇÃO LISTAFUNÇÕES",
		"boolean": "DECFUNÇÃO LISTAFUNÇÕES",
		"void":    "DECFUNÇÃO LISTAFUNÇÕES",
		"main":    "e",
		"$":       "e",
	},
	"DECFUNÇÃO": {
		"int":     "TIPORETORNO id ( PARÂMETROS ) BLOCO",
		"float":   "TIPORETORNO id ( PARÂMETROS ) BLOCO",
		"char":    "TIPORETORNO id ( PARÂMETROS ) BLOCO",
		"boolean": "TIPORETORNO id ( PARÂMETROS ) BLOCO",
		"void":    "TIPORETORNO id ( PARÂMETROS ) BLOCO",
		"main":    "TIPORETORNO id ( ) BLOCO",
	},
	"TIPORETORNO": {
		"int":     "TIPO",
		"float":   "TIPO",
		"char":    "TIPO",
		"boolean": "TIPO",
		"void":    "void",
	},
	"TIPO": {
		"int":     "TIPOBASE",
		"float":   "TIPOBASE",
		"char":    "TIPOBASE",
		"boolean": "TIPOBASE",
	},
	"TIPOBASE": {
		"int":     "int",
		"float":   "float",
		"char":    "char",
		"boolean": "boolean",
	},
	"DIMENSÃO": {
		"[":  "[ num_int ] DIMENSÃO",
		")":  "e",
		"(":  "e",
		"id": "e",
		",":  "e",
	},
	"PARÂMETROS": {
		"int":     "TIPO id LISTAPARAMETROS",
		"float":   "TIPO id LISTAPARAMETROS",
		"char":    "TIPO id LISTAPARAMETROS",
		"boolean": "TIPO id LISTAPARAMETROS",
		")":       "e",
	},
	"LISTAPARAMETROS": {
		")": "e",
		",": ", TIPO id LISTAPARAMETROS",
	},
	"PRINCIPAL": {
		"main": "main ( ) BLOCO",
	},
	"BLOCO": {
		"{": "{ COMANDOS }",
	},
	"LISTAVARIAVEIS": {
		"int":     "TIPO id LISTAID ; LISTAVARIAVEIS",
		"float":   "TIPO id LISTAID ; LISTAVARIAVEIS",
		"char":    "TIPO id LISTAID ; LISTAVARIAVEIS",
		"boolean": "TIPO id LISTAID ; LISTAVARIAVEIS",
		"}":       "e",
		"id":      "e",
		"=":       "e",
	},
	"LISTAID": {
		",": ", id LISTAID",
		";": "e",
	},
	"COMANDOS": {
		"int":     "COMANDO COMANDOS",
		"scanf":   "COMANDO COMANDOS",
		"println": "COMANDO COMANDOS",
		"id":      "COMANDO COMANDOS",
		"if":      "COMANDO COMANDOS",
		"while":   "COMANDO COMANDOS",
		"for":     "COMANDO COMANDOS",
		"return":  "COMANDO COMANDOS",
		"}":       "e",
		";":       "e",
	},
	"COMANDO": {
		"int":     "DECLARAÇÃO",
		"scanf":   "LEITURA",
		"println": "ESCRITA",
		"id":      "ATRIBUIÇÃO",
		"if":      "SELEÇÃO",
		"while":   "ENQUANTO",
		"return":  "RETORNO",
	},
	"LEITURA": {
		"scanf": "scanf ( TERMOLEITURA NOVOTERMOLEITURA ) ;",
	},
	"TERMOLEITURA": {
		"id": "id DIMENSAO2",
	},
	"NOVOTERMOLEITURA": {
		",": ", TERMOLEITURA NOVOTERMOLEITURA",
		")": "e",
	},
	"DIMENSAO2": {
		"[": "[ EXPR_ADITIVA ] DIMENSAO2",
		",": "e",
		")": "e",
	},
	"ESCRITA": {
		"println": "println ( TERMOESCRITA NOVOTERMOESCRITA ) ;",
	},
	"TERMOESCRITA": {
		"id":    "id DIMENSAO2",
		"texto": "texto",
	},
	"NOVOTERMOESCRITA": {
		",": ", TERMOESCRITA NOVOTERMOESCRITA",
		")": "e",
	},
	"SELEÇÃO": {
		"if": "if ( EXPRESSÃO ) BLOCO SENÃO",
	},
	"SENÃO": {
		"else": "else BLOCO",
		"}":    "e",
	},
	"ENQUANTO": {
		"while": "while ( EXPRESSÃO ) BLOCO",
	},
	"ATRIBUIÇÃO": {
		"id": "id = COMPLEMENTO ;",
	},
	"COMPLEMENTO": {
		"id":   "EXPRESSÃO",
		"func": "FUNÇÃO",
	},
	"FUNÇÃO": {
		"id": "func id ( ARGUMENTOS )",
	},
	"ARGUMENTOS": {
		"id": "EXPRESSÃO NOVO_ARGUMENTO",
		")":  "e",
	},
	"NOVO_ARGUMENTO": {
		",": ", EXPRESSÃO NOVO_ARGUMENTO",
		")": "e",
	},
	"RETORNO": {
		"return": "return EXPRESSAO ;",
	},
	"EXPRESSÃO": {
		"id":      "EXPR_OU",
		"num_int": "EXPR_OU",
		"num_dec": "EXPR_OU",
		"(":       "EXPR_OU",
		"+":       "EXPR_OU",
		"-":       "EXPR_OU",
		"!":       "EXPR_OU",
	},
	"EXPR_OU": {
		"id":      "EXPR_E EXPR_OU2",
		"num_int": "EXPR_E EXPR_OU2",
		"num_dec": "EXPR_E EXPR_OU2",
		"(":       "EXPR_E EXPR_OU2",
	},
	"EXPR_OU2": {
		"||": "|| EXPR_E EXPR_OU2",
		")":  "e",
	},
	"EXPR_E": {
		"id":      "EXPR_RELACIONAL EXPR_E2",
		"num_int": "EXPR_RELACIONAL EXPR_E2",
		"num_dec": "EXPR_RELACIONAL EXPR_E2",
		"(":       "EXPR_RELACIONAL EXPR_E2",
	},
	"EXPR_E2": {
		"&&": "&& EXPR_RELACIONAL EXPR_E2",
		")":  "e",
	},
	"EXPR_RELACIONAL": {
		"id":      "EXPR_ADITIVA EXPR_RELACIONAL2",
		"num_int": "EXPR_ADITIVA EXPR_RELACIONAL2",
		"num_dec": "EXPR_ADITIVA EXPR_RELACIONAL2",
		"(":       "EXPR_ADITIVA EXPR_RELACIONAL2",
	},
	"EXPR_RELACIONAL2": {
		"comp": "comp EXPR_ADITIVA",
		")":    "e",
	},
	"EXPR_ADITIVA": {
		"id":  "EXPR_MULTIPLICATIVA EXPR_ADITIVA2",
		"num": "EXPR_MULTIPLICATIVA EXPR_ADITIVA2",
		"(":   "EXPR_MULTIPLICATIVA EXPR_ADITIVA2",
	},
	"EXPR_ADITIVA2": {
		"+": "+ EXPR_MULTIPLICATIVA EXPR_ADITIVA2",
		"-": "- EXPR_MULTIPLICATIVA EXPR_ADITIVA2",
		")": "e",
	},
	"EXPR_MULTIPLICATIVA": {
		"id":  "FATOR EXPR_MULTIPLICATIVA2",
		"num": "FATOR EXPR_MULTIPLICATIVA2",
		"(":   "FATOR EXPR_MULTIPLICATIVA2",
	},
	"EXPR_MULTIPLICATIVA2": {
		"*": "* FATOR EXPR_MULTIPLICATIVA2",
		"/": "/ FATOR EXPR_MULTIPLICATIVA2",
		"%": "% FATOR EXPR_MULTIPLICATIVA2",
		")": "e",
	},
	"FATOR": {
		"id":      "TERMO",
		"num_int": "CONSTANTE",
		"num_dec": "CONSTANTE",
		"(":       "( EXPRESSÃO )",
		"!":       "! FATOR",
		"+":       "+ TERMO",
		"-":       "- TERMO",
	},
	"TERMO": {
		"id":  "id DIMENSAO2",
		"num": "CONSTANTE",
	},
	"CONSTANTE": {
		"num_int": "num_int",
		"num_dec": "num_dec",
	},
	"DECLARAÇÃO": {
		"int":    "TIPO id DECLARAÇÃO2",
		"bool":   "TIPO id DECLARAÇÃO2",
		"string": "TIPO id DECLARAÇÃO2",
		"float":  "TIPO id DECLARAÇÃO2",
	},
	"DECLARAÇÃO2": {
		";": "e",
	},
}

func Parse(tokens []lexical.Token) {
	stack := []string{"$", "PROGRAMA"}
	tokenIndex := 0

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		token := tokens[tokenIndex]

		fmt.Printf("Topo: %s, Token: %s\n", top, token.Value)

		if top == "$" && token.Type == lexical.EOF {
			fmt.Println("Parse completo com sucesso!")
			return
		}

		if isTerminal(top) {
			//fmt.Println(top, token.Value)
			if top == "id" && token.Type == lexical.ID {
				stack = stack[:len(stack)-1]
				tokenIndex++
			} else if top == "num_int" && token.Type == lexical.NUMBER && !strings.Contains(token.Value, ".") {
				stack = stack[:len(stack)-1]
				tokenIndex++
			} else if top == "num_dec" && token.Type == lexical.NUMBER && strings.Contains(token.Value, ".") {
				stack = stack[:len(stack)-1]
				tokenIndex++
			} else if top == token.Value {
				stack = stack[:len(stack)-1]
				tokenIndex++
			} else {
				fmt.Printf("Erro de sintaxe: esperado '%s', encontrado '%s'\n", top, token.Value)
				return
			}
		} else {
			//fmt.Println(top)
			production, exists := parsingTable[top][token.Value]
			//fmt.Println(production, exists)
			if !exists {
				fmt.Printf("Erro de sintaxe: nenhuma produção para o não-terminal '%s' com o token '%s'\n", top, token.Value)
				return
			}

			stack = stack[:len(stack)-1]
			if production != "e" {
				productions := strings.Split(production, " ")
				for i := len(productions) - 1; i >= 0; i-- {
					stack = append(stack, productions[i])
				}
			}
		}
	}

	fmt.Println("Erro de sintaxe: fim inesperado de entrada")
}

func isTerminal(symbol string) bool {
	terminals := map[string]bool{
		"main": true, "int": true, "char": true, "float": true, "boolean": true,
		"if": true, "for": true, "while": true, "return": true, "num_int": true, "num_dec": true,
		"(": true, ")": true, "{": true, "}": true, ";": true, "=": true, "id": true,
	}
	return terminals[symbol]
}
