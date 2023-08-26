package main // https://stepik.org/lesson/359395/step/4

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()

	// time.DateTime
	const layout = "2006-01-02 15:04:05"

	date, _ := time.Parse(layout, in.Text())

	if date.Hour() >= 13 {
		date = date.AddDate(0, 0, 1)
	}

	fmt.Println(date.Format(layout))
}
