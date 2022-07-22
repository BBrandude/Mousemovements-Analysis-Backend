package calculations

import (
	"math"
)

func CalculateStandardDeviation(mouseMovements []struct {
	X int     "json:\"x\""
	Y int     "json:\"y\""
	T float64 "json:\"t\""
}) float64 {
	movementsLength := len(mouseMovements)
	x1 := float64(mouseMovements[0].X)
	y1 := float64(mouseMovements[0].Y)
	x2 := float64(mouseMovements[movementsLength-1].X)
	y2 := float64(mouseMovements[movementsLength-1].Y)
	slope := (y2 - y1) / (x2 - x1)
	yIntercept := y1 - (slope * x1)
	var totalDistance float64 = 0
	for i := 0; i < movementsLength; i++ {
		totalDistance += shortestDistance(float64(mouseMovements[i].X), float64(mouseMovements[i].Y), -slope, 1, -yIntercept)
	}
	standardDeviation := totalDistance / float64(movementsLength)
	return standardDeviation
}

func shortestDistance(
	x1 float64,
	y1 float64,
	a float64,
	b float64,
	c float64,
) float64 {
	d := math.Abs(float64(a*x1+b*y1+c)) / math.Sqrt(float64(a*a+b*b))
	return d
}
