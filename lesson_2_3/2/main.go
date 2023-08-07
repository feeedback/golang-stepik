package main // https://stepik.org/lesson/266497/step/6

import "fmt"

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}
