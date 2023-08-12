package main // https://stepik.org/lesson/230630/step/11

import (
	"fmt"
	"strings"
)

func exercises5(str string) string {
	result := ""

	for _, char := range str {
		if strings.Count(str, string(char)) == 1 {
			result += string(char)
		}
	}
	return result
}

func main() {
	str := ""
	fmt.Scan(&str)

	fmt.Println(exercises5(str))
}
