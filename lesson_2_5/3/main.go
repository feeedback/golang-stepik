package main // https://stepik.org/lesson/230630/step/9

import (
	"fmt"
	"strings"
)

func exercises3(x, s string) int {
	index := strings.Index(x, s)

	if index == -1 {
		return -1
	} else {
		return index
	}
}

func main() {
	x, s := "", ""
	fmt.Scan(&x, &s)

	fmt.Println(exercises3(x, s))
}
