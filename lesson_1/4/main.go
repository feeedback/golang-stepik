package main // https://stepik.org/lesson/228261/step/12

import "fmt"

func arithmeticCalculations_4(a int) int {
	return a % 10
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(arithmeticCalculations_4(a))
}
