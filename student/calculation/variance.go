package formulas

import "math"

func CalculateVariance(data []float64, average float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(value-average, 2)
	}
	return sum / float64(len(data))
}
