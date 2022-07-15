package calculations

func CalcDistanceOverTime(distances []float64, times []float64) []float64 {

	distancesOverTimes := make([]float64, 0)
	for i := 0; i < len(distances); i++ {
		distancesOverTimes = append(distancesOverTimes, (distances[i] / times[i]))
	}
	return distancesOverTimes
}
