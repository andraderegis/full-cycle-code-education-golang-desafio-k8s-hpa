package arithmetics

import (
	"math"
)

func SumSqrtLoop(loop int, value int) float64 {
	i := 0
	sumSqrt := 0.0

	for i <= loop {
		i = i + 1
		sumSqrt = sumSqrt + math.Sqrt(float64(value))
	}

	return sumSqrt
}
