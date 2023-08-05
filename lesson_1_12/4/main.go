// https://stepik.org/lesson/228265/step/15
package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	// arr := make([]int, N, 100)
	var arr [100]int

	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
}
