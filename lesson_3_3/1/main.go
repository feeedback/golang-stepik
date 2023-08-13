package main // https://stepik.org/lesson/349644/step/7

import (
	"fmt"
	"strconv"
)

func main() {
	// code for stepik
	fn := func(n uint) uint {
		numberStr := strconv.Itoa(int(n))
		newNumberStr := ""

		for _, char := range numberStr {
			digit, _ := strconv.Atoi(string(char))

			if digit%2 == 0 && digit != 0 {
				newNumberStr += string(char)
			}
		}
		newNumber, _ := strconv.Atoi(newNumberStr)

		if newNumber == 0 {
			return 100
		}
		return uint(newNumber)
	}
	// code for stepik ------------

	fmt.Printf("fn %T", fn)
}
