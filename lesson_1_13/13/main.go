package main // https://stepik.org/lesson/229320/step/14

import "fmt"

func exercises13(N int) int {
	for count, a, b := 2, 0, 1; a+b <= N; a, b, count = b, a+b, count+1 {
		if a+b == N {
			return count
		}
	}
	return -1
}

func main() {
	var a int // натуральное число A > 1
	fmt.Scan(&a)

	fmt.Println(exercises13(a))
}
