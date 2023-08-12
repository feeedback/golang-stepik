package main // https://stepik.org/lesson/230630/step/7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func exercises1(text string) string {
	if len(text) > 0 && unicode.IsUpper([]rune(text)[0]) && strings.HasSuffix(text, ".") {
		return "Right"
	}
	return "Wrong"
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(exercises1(text))
}
