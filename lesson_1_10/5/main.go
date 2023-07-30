package main // https://stepik.org/lesson/228263/step/7

import (
	"fmt"
)

// первое число от 1 до n включительно, кратное c, но НЕ кратное d
func cycles_5(n, c, d int) int {
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			return i
		}
	}
	return -1
}

func main() {
	var a, b, c int // 3 натуральных числа n, c, d, каждое из которых не превышает 10000.
	fmt.Scan(&a, &b, &c)

	res := cycles_5(a, b, c)
	if res == -1 {
		return
	}
	fmt.Println(res)
}
