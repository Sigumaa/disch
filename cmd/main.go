package main

// disch for cli

// Receive text as an argument and display the converted text.

import (
	"bufio"
	"fmt"
	"github.com/Sigumaa/disch"
	"github.com/atotto/clipboard"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		s.Scan()
		input := s.Text()
		if input == "exit()" {
			break
		}
		res, err := disch.Convert(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err := clipboard.WriteAll(res); err != nil {
			fmt.Println(err)
			fmt.Println("Failed to copy to clipboard. Please copy manually.")
			fmt.Println(res)
			continue
		}
		fmt.Println("Copied!!")
	}
}
