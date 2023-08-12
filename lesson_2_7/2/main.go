package main // https://stepik.org/lesson/229321/step/3

import (
	"fmt"
	"strings"
)

func exercises2(a string) string {
	return strings.Join(strings.Split(a, ""), "*")
}

func main() {
	var a string
	fmt.Scan(&a)

	fmt.Println(exercises2(a))
}
