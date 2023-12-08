package internal

import (
	"fmt"
	"math"
)

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(GCD(a, b)))
}

func Quadratic(a, b, c float64) (float64, float64, error) {
	discriminant := (b * b) - (4 * a * c)
	rooted := math.Sqrt(discriminant)
	if math.IsNaN(rooted) {
		return math.NaN(), math.NaN(), fmt.Errorf("discriinant is less than zero: %v^2 - 4(%v)(%v) = %v", b, a, c, discriminant)
	}
	x1 := (-b + rooted) / (2 * a)
	x2 := (-b - rooted) / (2 * a)
	return x1, x2, nil
}
