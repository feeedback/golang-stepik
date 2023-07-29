package main // https://stepik.org/lesson/228261/step/11

import "fmt"

func arithmeticCalculations(a int) int {
	return a*2 + 100
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(arithmeticCalculations(a))
}
