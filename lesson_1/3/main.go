package main // https://stepik.org/lesson/228261/step/13

import "fmt"

func arithmeticCalculations_3(a int) int {
	return a * a
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(arithmeticCalculations_3(a))
}
