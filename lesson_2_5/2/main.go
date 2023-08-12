package main // https://stepik.org/lesson/230630/step/8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func exercises2(text string) string {
	runes := []rune(strings.ToLower(text))

	isPalindrome := true

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		return "Палиндром"
	} else {
		return "Нет"
	}
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Println(exercises2(text))
}
