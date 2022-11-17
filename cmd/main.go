package main

// disch for cli

// Receive text as an argument and display the converted text.

import (
	"fmt"
	"github.com/Sigumaa/disch"
)

func main() {
	for {
		fmt.Print("> ")
		var input string
		fmt.Scan(&input)
		if input == "exit()" {
			break
		}
		res, err := disch.Convert(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res)
	}
}
