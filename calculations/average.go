package calculations

func CalculateAverage(values []float64) float64 {
	valuesLength := len(values)
	var sum float64 = 0
	for i := 0; i < valuesLength; i++ {
		sum += values[i]
	}
	average := sum / float64(valuesLength)
	return average
}
