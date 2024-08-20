package input

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestProcessInput(t *testing.T) {
	// Define test cases
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          "1.0\n2.0\n",
			expectedOutput: "0 3\n",
		},
		{
			input:          "10\n20\n30\n",
			expectedOutput: "5 26\n3 37\n",
		},
		{
			input:          "5\n5\n5\n",
			expectedOutput: "5 5\n5 5\n",
		},
		{
			input:          "invalid\n10\n",
			expectedOutput: "Invalid input: strconv.ParseFloat: parsing \"invalid\": invalid syntax\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// Create a scanner from the input string
			scanner := bufio.NewScanner(strings.NewReader(tt.input))

			// Capture the output
			output := captureOutput(func() {
				ProcessInput(scanner)
			})

			if output != tt.expectedOutput {
				t.Errorf("Expected %q, but got %q", tt.expectedOutput, output)
			}
		})
	}
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	// Run the function while capturing its output
	f()

	// Close the write end of the pipe to stop capturing
	w.Close()

	// Restore the original stdout
	os.Stdout = stdout

	// Read the captured output from the read end of the pipe
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
