package main // https://stepik.org/lesson/230630/step/10

import "fmt"

func exercises4(str string) string {
	var oddChars string

	for i, char := range str {
		if i%2 == 1 {
			oddChars += string(char)
		}
	}
	return oddChars
}

func main() {
	str := ""
	fmt.Scan(&str)

	fmt.Println(exercises4(str))
}
