package main // https://stepik.org/lesson/350788/step/13

import (
	"fmt"
	"strings"
)

type Battery struct {
	Capacity string
}

func (b Battery) String() string {
	chargedCount := strings.Count(b.Capacity, "1")

	chargedPart := strings.Repeat("X", chargedCount)

	emptyPart := strings.Repeat(" ", len(b.Capacity)-chargedCount)

	return fmt.Sprintf("[%s%s]", emptyPart, chargedPart)
}

func main() {
	var input string
	fmt.Scan(&input)

	batteryForTest := Battery{Capacity: input}

	// for stepik only above
	fmt.Println(batteryForTest)
}
