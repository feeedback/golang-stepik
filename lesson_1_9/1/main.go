package main // https://stepik.org/lesson/232593/step/2

import "fmt"

func conditionalConstructs(a int) string {
	if a > 0 {
		return "Число положительное"
	} else if a < 0 {
		return "Число отрицательное"
	} else {
		return "Ноль"
	}
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(conditionalConstructs(a))
}
