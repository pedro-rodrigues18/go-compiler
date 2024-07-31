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
	file, err := os.Open("./input/main.txt")
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
	tokens := lexical.Scan(sourceCode)

	for _, token := range tokens {
		fmt.Printf("%v: %v\n", token.Type, token.Value)
	}
}
