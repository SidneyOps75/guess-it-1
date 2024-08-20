package main

import (
	"bufio"
	"fmt"
	"os"

	input "guess-it-1/processInput"
)

func main() {
	// Check if the command-line argument count is correct.
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}
	// Scan the standard Input and return the lower and the upper range using the ProcessInput function.
	scanner := bufio.NewScanner(os.Stdin)
	input.ProcessInput(scanner)
}
