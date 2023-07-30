package main // https://stepik.org/lesson/232593/step/8

import (
	"fmt"
	"math"
)

func digitize(n, ordinal int) int {
	return (n / int(math.Pow(10, float64(ordinal-1)))) % 10
}

func conditionalConstructs_4(n int) string {
	var a, b, c, d, e, f int

	a = digitize(n, 1)
	b = digitize(n, 2)
	c = digitize(n, 3)
	d = digitize(n, 4)
	e = digitize(n, 5)
	f = digitize(n, 6)

	if a+b+c == d+e+f {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var a int // одно шестизначное число
	fmt.Scan(&a)

	fmt.Println(conditionalConstructs_4(a))
}
