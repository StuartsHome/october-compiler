package main

import (
	"fmt"

	"main.go/service"
)

func main() {
	fmt.Println("Hello world")

	var Input = "LET what foobar = 123"
	lexObj := service.NewLex(Input)

	for lexObj.Peek(Input) != `\0` {
		fmt.Println(lexObj.CurrChar)
		lexObj.NextChar(Input)
	}

}
