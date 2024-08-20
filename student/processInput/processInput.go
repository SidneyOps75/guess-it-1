package input

import (
	"bufio"
	"fmt"
	"strconv"

	formulas "guess-it-1/calculation"
)

// ProcessInput processes input data from the scanner and performs calculations.
func ProcessInput(scanner *bufio.Scanner) {
	// Declares a slice of float64 to store the input numbers.
	var data []float64

	// Starts a loop to scan the input line by line.
	for scanner.Scan() {
		// Converts the scanned text into a float64.
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue // Skips to the next iteration of the loop if the input is invalid.
		}

		// Adds the valid float number to the data slice.
		data = append(data, num)
		// Checks if the data slice contains more than one number.
		if len(data) > 1 {
			// Calls the Range function from the calc package, passing the data slice.
			lower, upper := formulas.Range(data)
			// Prints the calculated lower and upper range values.
			fmt.Printf("%d %d\n", lower, upper)
		}
	}

	// Checks if there was any error during the scanning process.
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
}
