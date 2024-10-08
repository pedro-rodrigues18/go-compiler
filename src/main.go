/*
**	Authors: Pedro Henrique Rodrigues Pereira;
			 Gabriel Henrique Silva Gontijo
**	Date: 27/04/2024
**	Update: 08/09/2024
**/

package main

import (
	"fmt"
	"os"

	"github.com/pedro-rodrigues18/go-compiler/src/lexical"
	"github.com/pedro-rodrigues18/go-compiler/src/parser"
	"github.com/pedro-rodrigues18/go-compiler/src/syntax"

	"github.com/antlr4-go/antlr/v4"
)

type MyErrorListener struct {
	*antlr.DefaultErrorListener
}

func (d *MyErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Printf("Erro de sintaxe na linha %d:%d - %s\n", line, column, msg)
}

type myListener struct {
	*parser.BaseGrammarListener
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])

	lexer := parser.NewGrammarLexer(input)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	fmt.Println("Tokens:")
	tokens.Fill()
	for _, token := range tokens.GetAllTokens() {
		tokenType := token.GetTokenType()
		if tokenType == antlr.TokenEOF {
			fmt.Println("EOF")
		} else if tokenType == -1 {
			fmt.Printf("Token não reconhecido: %s\n", token.GetText())
		} else {
			tokenName := lexer.SymbolicNames[tokenType]
			if tokenName == "" {
				tokenName = "UNKNOWN"
			}
			fmt.Printf("%s (%s)\n", tokenName, token.GetText())
		}
	}

	fmt.Print("\n")

	p := parser.NewGrammarParser(tokens)

	errorListener := &MyErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.Programa()

	fmt.Println("\nÁrvore sintática:")
	fmt.Println(tree.ToStringTree(nil, p))

	listener := &myListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	fmt.Println("\nAnálise sintática com ANTLR concluída.")

	fmt.Println("=====================================")

	file, err := os.Open("./input/teste.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	size := fileInfo.Size()
	fileContent := make([]byte, size)

	_, err = file.Read(fileContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	sourceCode := string(fileContent)
	newTokens := lexical.Scan(sourceCode)

	fmt.Println("Tokens:")
	for _, token := range newTokens {
		fmt.Printf("%v: %v\n", token.Type, token.Value)
	}

	fmt.Println("\nAnálise Sintática Manual:")

	syntax.Parse(newTokens)

}
