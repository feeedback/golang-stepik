package main // https://stepik.org/lesson/228838/step/11

import "fmt"

func sumInt(nums ...int) (int, int) {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return len(nums), sum
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(sumInt(a))
}
