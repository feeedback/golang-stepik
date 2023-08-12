package main // https://stepik.org/lesson/229321/step/4

import (
	"fmt"
	"strconv"
)

func exercises3(input string) int {
	maxDigit := 0

	for _, char := range input {
		digit, _ := strconv.Atoi(string(char))

		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}

func main() {
	var a string
	fmt.Scan(&a)

	fmt.Println(exercises3(a))
}
