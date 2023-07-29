package main // https://stepik.org/lesson/228261/step/12

import "fmt"

func arithmeticCalculations_2(a int, b int) int {
	return a*a + b*b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(arithmeticCalculations_2(a, b))
}
