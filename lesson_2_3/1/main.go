package main // https://stepik.org/lesson/266497/step/5

import "fmt"

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}
