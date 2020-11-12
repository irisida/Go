package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	p := fmt.Println

	scanner := bufio.NewScanner(os.Stdin)
	text := scanner.Text()

	for scanner.Scan() {
		text = scanner.Text()
		p(text)

		if text == "exit" {
			p("exiting the scanner!")
			break
		}
	}

	p("Bye!")
}
