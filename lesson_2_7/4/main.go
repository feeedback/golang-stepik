package main // https://stepik.org/lesson/229321/step/5

import (
	"fmt"
	"strconv"
)

func exercises4(numStr string) string {
	resStr := ""

	for i := 0; i < len(numStr); i++ {
		digit, _ := strconv.Atoi(string(numStr[i]))

		resStr += strconv.Itoa(digit * digit)
	}
	return resStr
}

func main() {
	var a string
	fmt.Scan(&a)

	fmt.Println(exercises4(a))
}
