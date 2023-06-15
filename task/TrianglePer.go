package task

import "math"

func Triangle(a, b float64) float64 {
	c := math.Sqrt((a * a) + (b * b))
	Perimetr := a + b + c
	return Perimetr
}
