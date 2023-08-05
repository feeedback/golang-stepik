// https://stepik.org/lesson/229320/step/7
package main

import "fmt"

func exercises6(a, b int) float64 {
	average := float64(a+b) / 2
	return average
}

func main() {
	var a, b int // два целых положительных числа a и b.
	fmt.Scan(&a, &b)

	fmt.Println(exercises6(a, b))
}
