package main // https://stepik.org/lesson/229320/step/10

import "fmt"

func calculateDigitalRoot(n int) int {
	for n >= 10 {
		sum := 0

		for n > 0 {
			sum += n % 10
			n /= 10
		}

		n = sum
	}
	return n
	// digitalRoot := (n - 1) % 9 + 1
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(calculateDigitalRoot(a))
}
