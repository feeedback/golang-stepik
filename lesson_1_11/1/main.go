// https://stepik.org/lesson/351787/step/3
package main

import (
	"fmt"
	"math"
)

func stringFormat_1(num float64) string {
	if num <= 0 {
		return fmt.Sprintf("число %.2f не подходит", num)
	}
	if num > 10000 {
		return fmt.Sprintf("%e", num)
	}

	squaredNumber := math.Pow(num, 2)
	truncatedNumber := math.Round(squaredNumber*10000) / 10000

	return fmt.Sprintf("%.4f", truncatedNumber)
}

func main() {
	var n float64

	fmt.Scan(&n)

	fmt.Println(stringFormat_1(n))
}
