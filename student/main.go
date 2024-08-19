package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	calc "guess-it-1/calculation"
)

func main() {
	// Ensure no command-line arguments are provided
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}

	var data []float64
	scanner := bufio.NewScanner(os.Stdin)

	// Continuously scan input until EOF or error
	for scanner.Scan() {
		// Parse input as float64
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		// Append to data and calculate range if more than one data point
		data = append(data, num)
		if len(data) > 1 {
			lower, upper := calc.Range(data)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}

	// Handle potential scanning errors
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
}
