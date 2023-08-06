package main // https://stepik.org/lesson/229320/step/16

import "fmt"

func exercises15(num, queryRemoveDigit int) int {
	filteredNum := 0

	for multi := 1; num > 0; {
		currentDigit := num % 10
		num /= 10

		if currentDigit != queryRemoveDigit {
			filteredNum += multi * currentDigit
			multi *= 10
		}
	}
	return filteredNum
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(exercises15(a, b))
}
