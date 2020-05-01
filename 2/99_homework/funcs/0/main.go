package main

import (
	"fmt"
	"math"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	if (x <= 0) {
		return 0
	}

	var prev float64 = 0
	var curr float64 =  x / 2// Initial guess

	for math.Abs(curr - prev) > 0.01 {
		prev = curr
		curr = 0.5 * (prev + x / prev);
	}

	return curr
}

func main() {
	fmt.Println(Sqrt(1))
}
