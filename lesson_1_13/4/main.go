// https://stepik.org/lesson/229320/step/5
package main

import (
	"fmt"
	"math"
)

func exercises4(a, b, c float64) string {
	maxSide := math.Max(a, math.Max(b, c))

	isRight := false

	switch maxSide {
	case a:
		isRight = a*a == b*b+c*c
	case b:
		isRight = b*b == a*a+c*c
	case c:
		isRight = c*c == a*a+b*b
	}

	if isRight {
		return "Прямоугольный"
	}
	return "Непрямоугольный"

}

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	fmt.Println(exercises4(a, b, c))
}
