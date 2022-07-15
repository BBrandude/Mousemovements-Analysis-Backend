package calculations

import (
	"math"
)

func CalculateDistance(x1 int, x2 int, y1 int, y2 int) float64 {
	xdifference := x2 - x1
	ydifference := y2 - y1
	inside := math.Pow(float64(xdifference), 2) + math.Pow(float64(ydifference), 2)
	distance := math.Sqrt(inside)
	return distance
}
