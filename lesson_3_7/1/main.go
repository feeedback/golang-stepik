package main // https://stepik.org/lesson/359395/step/3

import (
	"fmt"
	"time"
)

func main() {
	var input string
	fmt.Scan(&input)

	t, _ := time.Parse(time.RFC3339, input)

	outputFormat := time.UnixDate

	fmt.Println(t.Format(outputFormat))
}
