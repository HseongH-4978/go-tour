package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x <= 0 {
		return 0
	}

	z := float64(1)

	for y := float64(0); math.Abs(z-y) >= 1e-9; {
		y = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
