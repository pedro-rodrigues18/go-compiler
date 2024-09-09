package syntax

import (
	"fmt"
	"strings"

	"github.com/pedro-rodrigues18/go-compiler/src/lexical"
)

var parsingTable = map[string]map[string]string{
	"programa": {
		"main": "listaFuncoes principal EOF",
	},
	"listaFuncoes": {
		"int":     "decFuncao listaFuncoes",
		"float":   "decFuncao listaFuncoes",
		"char":    "decFuncao listaFuncoes",
		"boolean": "decFuncao listaFuncoes",
		"void":    "decFuncao listaFuncoes",
		"main":    "",
		"$":       "",
	},
	"decFuncao": {
		"int":     "tipoRetorno ID ( parametros ) bloco",
		"float":   "tipoRetorno ID ( parametros ) bloco",
		"char":    "tipoRetorno ID ( parametros ) bloco",
		"boolean": "tipoRetorno ID ( parametros ) bloco",
		"void":    "tipoRetorno ID ( parametros ) bloco",
	},
	"tipoRetorno": {
		"int":     "tipo",
		"float":   "tipo",
		"char":    "tipo",
		"boolean": "tipo",
		"void":    "void",
	},
	"tipo": {
		"int":     "tipoBase dimensao",
		"float":   "tipoBase dimensao",
		"char":    "tipoBase dimensao",
		"boolean": "tipoBase dimensao",
	},
	"tipoBase": {
		"int":     "int",
		"float":   "float",
		"char":    "char",
		"boolean": "boolean",
	},
	"dimensao": {
		"[":  "[ NUM_INT ] dimensao",
		")":  "",
		"ID": "",
	},
	"parametros": {
		"int":     "tipo ID listaParametros",
		"float":   "tipo ID listaParametros",
		"char":    "tipo ID listaParametros",
		"boolean": "tipo ID listaParametros",
		")":       "",
	},
	"listaParametros": {
		",": ", tipo ID listaParametros",
		")": "",
	},
	"principal": {
		"main": "main ( ) bloco",
	},
	"bloco": {
		"{": "{ listaVariaveis comandos }",
	},
	"listaVariaveis": {
		"int":     "tipo ID listaId ; listaVariaveis",
		"float":   "tipo ID listaId ; listaVariaveis",
		"char":    "tipo ID listaId ; listaVariaveis",
		"boolean": "tipo ID listaId ; listaVariaveis",
		"return":  "",
		"}":       "",
	},
	"listaId": {
		",": ", ID listaId",
		";": "",
	},
	"comandos": {
		"scanf":   "comando comandos",
		"println": "comando comandos",
		"ID":      "comando comandos",
		"if":      "comando comandos",
		"while":   "comando comandos",
		"return":  "comando comandos",
		"}":       "",
	},
	"comando": {
		"scanf":   "leitura",
		"println": "escrita",
		"ID":      "atribuicao",
		"if":      "selecao",
		"while":   "enquanto",
		"return":  "retorno",
	},
	"leitura": {
		"scanf": "scanf ( termoLeitura novoTermoLeitura ) ;",
	},
	"termoLeitura": {
		"ID": "ID dimensao2",
	},
	"novoTermoLeitura": {
		",": ", termoLeitura novoTermoLeitura",
		")": "",
	},
	"dimensao2": {
		"[": "[ exprAditiva ] dimensao2",
		",": "",
		")": "",
	},
	"escrita": {
		"println": "println ( termoEscrita novoTermoEscrita ) ;",
	},
	"termoEscrita": {
		"ID":     "ID dimensao2",
		"NUMBER": "NUMBER",
		"TEXTO":  "TEXTO",
	},
	"novoTermoEscrita": {
		",": ", termoEscrita novoTermoEscrita",
		")": "",
	},
	"selecao": {
		"if": "if ( expressao ) bloco senao",
	},
	"senao": {
		"else": "else bloco",
		"}":    "",
	},
	"enquanto": {
		"while": "while ( expressao ) bloco",
	},
	"atribuicao": {
		"ID": "ID = complemento ;",
	},
	"complemento": {
		"ID":   "expressao",
		"func": "funcao",
	},
	"funcao": {
		"func": "func ID ( argumentos )",
	},
	"argumentos": {
		"ID":     "expressao novoArgumento",
		"NUMBER": "expressao novoArgumento",
		")":      "",
	},
	"novoArgumento": {
		",": ", expressao novoArgumento",
		")": "",
	},
	"retorno": {
		"return": "return expressao ;",
	},
	"expressao": {
		"ID":     "exprOu",
		"NUMBER": "exprOu",
		"(":      "exprOu",
		"!":      "exprOu",
		"+":      "exprOu",
		"-":      "exprOu",
	},
	"exprOu": {
		"ID":     "exprE exprOu2",
		"NUMBER": "exprE exprOu2",
		"(":      "exprE exprOu2",
		"!":      "exprE exprOu2",
		"+":      "exprE exprOu2",
		"-":      "exprE exprOu2",
	},
	"exprOu2": {
		"||": "|| exprE exprOu2",
		")":  "",
		";":  "",
	},
	"exprE": {
		"ID":     "exprRelacional exprE2",
		"NUMBER": "exprRelacional exprE2",
		"(":      "exprRelacional exprE2",
		"!":      "exprRelacional exprE2",
		"+":      "exprRelacional exprE2",
		"-":      "exprRelacional exprE2",
	},
	"exprE2": {
		"&&": "&& exprRelacional exprE2",
		")":  "",
		";":  "",
	},
	"exprRelacional": {
		"ID":     "exprAditiva exprRelacional2",
		"NUMBER": "exprAditiva exprRelacional2",
		"(":      "exprAditiva exprRelacional2",
		"!":      "exprAditiva exprRelacional2",
		"+":      "exprAditiva exprRelacional2",
		"-":      "exprAditiva exprRelacional2",
		";":      "",
	},
	"exprRelacional2": {
		"COMP": "COMP exprAditiva",
		")":    "",
		";":    "",
	},
	"exprAditiva": {
		"ID":     "exprMultiplicativa exprAditiva2",
		"NUMBER": "exprMultiplicativa exprAditiva2",
		"(":      "exprMultiplicativa exprAditiva2",
		"!":      "exprMultiplicativa exprAditiva2",
		"+":      "exprMultiplicativa exprAditiva2",
		"-":      "exprMultiplicativa exprAditiva2",
	},
	"exprAditiva2": {
		"+": "+ exprMultiplicativa exprAditiva2",
		"-": "- exprMultiplicativa exprAditiva2",
		")": "",
		";": "",
	},
	"exprMultiplicativa": {
		"ID":     "fator exprMultiplicativa2",
		"NUMBER": "fator exprMultiplicativa2",
		"(":      "fator exprMultiplicativa2",
		"!":      "fator exprMultiplicativa2",
		"+":      "fator exprMultiplicativa2",
		"-":      "fator exprMultiplicativa2",
		";":      "",
	},
	"exprMultiplicativa2": {
		"*": "* fator exprMultiplicativa2",
		"/": "/ fator exprMultiplicativa2",
		"%": "% fator exprMultiplicativa2",
		")": "",
		";": "",
	},
	"fator": {
		"ID":     "sinal termo",
		"NUMBER": "sinal termo",
		"!":      "! fator",
		"(":      "( expressao )",
		"TEXTO":  "TEXTO",
	},
	"termo": {
		"ID":     "ID dimensao2",
		"NUMBER": "NUMBER",
	},
	"sinal": {
		"+":      "+",
		"-":      "-",
		"ID":     "",
		"NUMBER": "",
		"(":      "",
		"":       "",
	},
}

func Parse(tokens []lexical.Token) {
	stack := []string{"$", "programa"}
	tokenIndex := 0

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		token := tokens[tokenIndex]

		fmt.Printf("Topo: %s, Token: %s\n", top, token.Value)

		if top == "EOF" && token.Type == lexical.EOF {
			fmt.Println("Parse completo com sucesso!")
			return
		}

		if isTerminal(top) {
			// Verificação de terminais
			if top == "ID" && token.Type == lexical.ID {
				stack = stack[:len(stack)-1]
				tokenIndex++
			} else if top == "NUMBER" && token.Type == lexical.NUMBER {
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
			// Não-terminal procurar a produção correspondente na tabela
			production, exists := parsingTable[top][convertToString(token.Type, token.Value)]

			if !exists {
				fmt.Printf("Erro de sintaxe: nenhuma produção para o não-terminal '%s' com o token '%s'\n", top, token.Value)
				return
			}

			stack = stack[:len(stack)-1]

			if production != "" {
				productions := strings.Split(production, " ")
				for i := len(productions) - 1; i >= 0; i-- {
					stack = append(stack, productions[i])
				}
			}
		}
	}

	fmt.Println("Erro de sintaxe: fim inesperado de entrada")
}

func convertToString(t lexical.TokenType, value string) string {
	switch t {
	case lexical.KEYWORD:
		return value // Retorna o valor literal
	case lexical.ID:
		return "ID"
	case lexical.NUMBER:
		return "NUMBER"
	case lexical.SYMBOL:
		return value
	case lexical.STRING:
		return "STRING"
	case lexical.CHARACTER:
		return "CHARACTER"
	case lexical.EOF:
		return "$"
	default:
		return ""
	}
}

func isTerminal(symbol string) bool {
	terminals := map[string]bool{
		"main": true, "int": true, "char": true, "float": true, "boolean": true,
		"if": true, "else": true, "while": true, "return": true,
		"NUMBER": true, "ID": true, "STRING": true,
		"(": true, ")": true, "{": true, "}": true, ";": true, "=": true,
	}
	return terminals[symbol]
}
