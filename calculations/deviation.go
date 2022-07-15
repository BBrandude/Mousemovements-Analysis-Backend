package calculations

import "fmt"

func CalculateDeviation(mouseMovements []struct {
	X int     "json:\"x\""
	Y int     "json:\"y\""
	T float64 "json:\"t\""
}) {
	movementsLength := len(mouseMovements)
	x1 := float64(mouseMovements[0].X)
	y1 := float64(mouseMovements[0].Y)
	x2 := float64(mouseMovements[movementsLength-1].X)
	y2 := float64(mouseMovements[movementsLength-1].Y)
	slope := (y2 - y1) / (x2 - x1)
	yIntercept := y1 - (slope * x1)
	fmt.Println(yIntercept)
}
