/*
**	Author: Pedro H. R. Pereira
**	Date: 27/04/2024
**/

package main

import (
	"fmt"
	"os"

	"github.com/pedro-rodrigues18/go-compiler/src/lexical"
)

func main() {
	// Read the input file
	file, _ := os.Open("./input/main.txt")

	fileInfo, _ := file.Stat()

	// fmt.Println(fileInfo.Size())

	var size int64 = fileInfo.Size()

	fileContent := make([]byte, size)

	_, err := file.Read(fileContent)

	if err != nil {
		fmt.Println(err)
	}

	sourceCode := string(fileContent)
	file.Close()

	// Call function Scan() to get all tokens
	tokens := lexical.Scan(sourceCode)

	// Print all tokens and their types
	for _, token := range tokens {
		fmt.Printf("%v: %v\n", token.Type, token.Value)
	}
}
