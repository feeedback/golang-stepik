package main // https://stepik.org/lesson/229320/step/15

import (
	"fmt"
	// "strconv"
)

func exercises14(a int) string {
	return fmt.Sprintf("%b", a)
}

func main() {
	var a int // натуральное число
	fmt.Scan(&a)

	// fmt.Println(strconv.FormatInt(int64(a), 2))
	// fmt.Printf("%b", a)
	fmt.Println(exercises14(a))
}
