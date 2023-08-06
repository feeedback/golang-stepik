package main // https://stepik.org/lesson/228838/step/9

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(fibonacci(a))
}
