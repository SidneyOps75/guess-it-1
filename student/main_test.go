package main

import (
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic input",
			input:    "10\n20\n30\n",
			expected: "10 50\n",
		},
		{
			name:     "Single input",
			input:    "10\n",
			expected: "",
		},
		{
			name:     "Negative and positive input",
			input:    "-10\n0\n10\n",
			expected: "-30 30\n",
		},
		{
			name:     "Fractional input",
			input:    "1.5\n2.5\n3.5\n",
			expected: "0 5\n",
		},
		{
			name:     "Mixed input",
			input:    "5\n-5\n10\n-10\n",
			expected: "-10 10\n",
		},
		{
			name:     "No input",
			input:    "",
			expected: "",
		},
		{
			name:     "Multiple inputs, all the same",
			input:    "7\n7\n7\n7\n7\n7\n7\n",
			expected: "7 7\n", // All values are the same, so range should be the same
		},
		{
			name:     "Multiple inputs, all zero",
			input:    "0\n0\n0\n0\n0\n",
			expected: "0 0\n", // All zero values
		},
		{
			name:     "Very large input",
			input:    "1000000000\n2000000000\n3000000000\n4000000000\n5000000000\n",
			expected: "0 6000000000\n", // Large numbers to test upper bounds
		},
		{
			name:     "Very small input",
			input:    "0.000001\n0.000002\n0.000003\n0.000004\n0.000005\n",
			expected: "-0 0\n", // Very small fractional numbers
		},
		{
			name:     "Non-numeric input",
			input:    "abc\n123\n",
			expected: "123 123\n", // Invalid input should be ignored, resulting in only the valid number
		},
		{
			name:     "Input with whitespace",
			input:    " 10\n \n20\n  \n30\n",
			expected: "10 50\n", // Whitespace should be ignored
		},
		{
			name:     "Large number of inputs",
			input:    strings.Repeat("1\n", 1000), // 1000 inputs of '1'
			expected: "1 1\n",                     // All values are the same
		},
		{
			name:     "Edge case with negative large numbers",
			input:    strings.Join([]string{"-1000000", "-2000000", "-3000000", "-4000000", "-5000000"}, "\n") + "\n",
			expected: "-5000000 -1000000\n", // Large negative numbers
		},
		{
			name:     "Invalid number followed by valid numbers",
			input:    "abc\n10\n20\n30\n",
			expected: "10 50\n", // Invalid input should be ignored
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
