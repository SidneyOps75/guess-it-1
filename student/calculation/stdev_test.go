package formulas

import (
	"testing"
)

func TestStandarddev(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		variance       float64
		expectedStdDev float64
	}{
		{
			name:           "Zero variance",
			variance:       0,
			expectedStdDev: 0,
		},
		{
			name:           "Positive variance",
			variance:       4,
			expectedStdDev: 2,
		},
		{
			name:           "Small variance",
			variance:       0.25,
			expectedStdDev: 0.5,
		},
		{
			name:           "Large variance",
			variance:       100,
			expectedStdDev: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Standarddev(tt.variance)
			if result != tt.expectedStdDev {
				t.Errorf("Standarddev(%f) = %f; want %f", tt.variance, result, tt.expectedStdDev)
			}
		})
	}
}
