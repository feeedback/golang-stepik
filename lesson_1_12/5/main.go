// https://stepik.org/lesson/228265/step/16
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	positiveCount := 0

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)

		if num > 0 {
			positiveCount++
		}
	}

	fmt.Println(positiveCount)
}
