package main

// disch for cli

// Receive text as an argument and display the converted text.

import (
	"fmt"
	"github.com/Sigumaa/disch"
	"github.com/atotto/clipboard"
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
		//fmt.Println(res)
		if err := clipboard.WriteAll(res); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Copied!!")
	}
}
