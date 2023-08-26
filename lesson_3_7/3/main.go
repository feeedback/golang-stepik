package main // https://stepik.org/lesson/359395/step/7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const layoutDateTime = "02.01.2006 15:04:05"

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	dateStrings := strings.Split(input, ",")

	date1, _ := time.Parse(layoutDateTime, dateStrings[0])
	date2, _ := time.Parse(layoutDateTime, dateStrings[1])

	var duration time.Duration

	if date1.After(date2) {
		duration = date1.Sub(date2)
	} else {
		duration = date2.Sub(date1)
	}

	fmt.Println(duration)
}
