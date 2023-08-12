package main // https://stepik.org/lesson/229321/step/2

import (
	"fmt"
	"math"
)

func calculateHypotenuse(a float64, b float64) float64 {
	// return math.Sqrt(a*a + b*b)
	return math.Hypot(a, b)
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	fmt.Println(calculateHypotenuse(a, b))
}
