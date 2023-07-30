package main // https://stepik.org/lesson/228263/step/10

import (
	"fmt"
	"strconv"
	"strings"
)

func getCommonDigits(a, b int) []string {
	digits1 := strconv.Itoa(a)
	digits2 := strconv.Itoa(b)

	var commonDigits []string

	for i := 0; i < len(digits1); i++ {
		for j := 0; j < len(digits2); j++ {
			if digits1[i] == digits2[j] {
				commonDigits = append(commonDigits, string(digits1[i]))
			}
		}
	}

	return commonDigits
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	common := getCommonDigits(a, b)

	fmt.Println(strings.Join(common, " "))
}
