package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x < 0.0 {
		return math.NaN()
	}
	if x == 0.0 || x == 1.0 {
		return x
	}
	z := x
	for eps, d := 1e-10, 1.0; eps <= d; {
		d = math.Abs(z * z - x) / 2.0 / z
		z = math.Abs(z - d)
		fmt.Println(z, d)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1.44))
}
