package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	res := squrt(x)
	fmt.Println("Square root of", x, "is", res)
}

func squrt(num float64) float64 {
	z := 1.0
	var prev float64
	steps := 0
	for i := 0; i < 10; i++ {
		prev = z
		z -= (z*z - num) / (2 * z)
		steps++
		fmt.Println(z)
		if math.Abs(z-prev) < 1e-10 {
			break
		}

	}
	fmt.Println("steps =", steps)
	return z
}
