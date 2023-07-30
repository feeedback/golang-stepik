package main // https://stepik.org/lesson/228263/step/3

import "fmt"

func cycles_2(from int, to int) int {
	sum := 0

	for i := from; i <= to; i++ {
		sum += i
	}
	return sum
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(cycles_2(a, b))
}
