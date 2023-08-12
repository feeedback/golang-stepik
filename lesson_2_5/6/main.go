package main // https://stepik.org/lesson/230630/step/12

import (
	"fmt"
	"unicode"
)

func exercises6(password string) string {
	if len(password) < 5 {
		return "Wrong password"
	}

	for _, char := range password {
		if !(unicode.Is(unicode.Latin, char) || unicode.IsDigit(char)) {
			return "Wrong password"
		}
	}
	return "Ok"
}

func main() {
	str := ""
	fmt.Scan(&str)

	fmt.Println(exercises6(str))
}
