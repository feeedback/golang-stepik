// https://stepik.org/lesson/229320/step/8
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	zeroCount := 0

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)

		if num == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}
