package formulas

import "math"

// Range calculates the lower and upper bounds based on the mean and
// standard deviation of the last few elements in the data slice.
func Range(data []float64) (int, int) {
	// Extract the last 4 elements (or fewer if data has less than 4 elements)
	sampleData := getLastElements(data, 4)

	// Calculate the mean (average) of the sample data
	mean := CalculateAverage(sampleData)

	// Calculate the variance of the sample data using the mean
	variance := CalculateVariance(sampleData, mean)

	// Calculate the standard deviation from the variance
	stddev := Standarddev(variance)

	// Calculate the lower bound by subtracting 2.1 times the standard deviation from the mean
	lower := math.Round(mean - (2.1 * stddev))

	// Calculate the upper bound by adding 2.1 times the standard deviation to the mean
	upper := math.Round(mean + (2.1 * stddev))

	// Return the lower and upper bounds as integers
	return int(lower), int(upper)
}

// getLastElements returns the last 'count' elements of the data slice.
// If the data has fewer elements than 'count', it returns the entire slice.
func getLastElements(data []float64, count int) []float64 {
	// Calculate the starting index for the last 'count' elements
	start := len(data) - count

	// Ensure the start index is not negative
	if start < 0 {
		start = 0
	}

	// Return the slice starting from the calculated index to the end
	return data[start:]
}
