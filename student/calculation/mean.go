package calc

func CalculateAverage(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	average := float64(sum) / float64(len(data))
	return average
}
