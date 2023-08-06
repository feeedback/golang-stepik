package main // https://stepik.org/lesson/228838/step/7

import (
	"fmt"
	"math"
)

func exercises_2_2(a, b, c, d int) int {
	return int(math.Min(float64(a), math.Min(float64(b), math.Min(float64(c), float64(d)))))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(exercises_2_2(a, b, c, d))
}

// ***** for stepik ****

// import "math"

// func exercises_2_2(a, b, c, d int) int {
// 	return int(math.Min(float64(a), math.Min(float64(b), math.Min(float64(c), float64(d)))))
// }

// func minimumFromFour() int {
// 	var a, b, c, d int
// 	fmt.Scan(&a, &b, &c, &d)

// 	return exercises_2_2(a, b, c, d)
// }
