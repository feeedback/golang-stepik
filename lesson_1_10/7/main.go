package main // https://stepik.org/lesson/228263/step/9

import "fmt"

// Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная
// часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько
// лет вклад составит не менее y рублей.
func cycles_7(x, p, y int) int {
	years := 0

	for current := x; current < y; current += p * current / 100 {
		years++

		if current >= y {
			break
		}
	}
	return years
}

func main() {
	var a, b, c int // три натуральных числа
	fmt.Scan(&a, &b, &c)

	fmt.Println(cycles_7(a, b, c))
}
