package main

import (
	"fmt"

	str "golang.org/x/example/stringutil"
)

func main() {
	inputWords := "Hello, OTUS!"
	fmt.Println(str.Reverse(inputWords))
}
