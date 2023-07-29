package main // https://stepik.org/lesson/228261/step/15

import "fmt"

func arithmeticCalculations_5(a int) int {
	return a / 10 % 10
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(arithmeticCalculations_5(a))
}
