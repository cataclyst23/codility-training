package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("One argument expected.")
		os.Exit(1)
	}
	fmt.Printf("Result: %d", Solution(os.Args[1]))
}

func Solution(S string) int {
	insertables := 0
	counter := 0

	for _, c := range S {
		if c == 'a' {
			if counter == 2 {
				insertables = -1
				break
			}
			if counter < 2 {
				counter++
			}
		} else {
			if counter <= 2 {
				insertables += 2 - counter
				counter = 0
			}
		}
	}

	// After the last character
	insertables += 2 - counter

	return insertables
}
