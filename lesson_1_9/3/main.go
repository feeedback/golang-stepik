package main // https://stepik.org/lesson/232593/step/7

import "fmt"

func conditionalConstructs_3(a int) int {
	if a < 10 {
		return a
	} else if a > 9 && a < 100 {
		return a / 10
	} else if a > 99 && a < 1000 {
		return a / 100
	} else if a > 990 && a < 10000 {
		return a / 1000
	} else {
		return a / 10000
	}

}

func main() {
	var a int // натуральное число, не превосходящее 10000
	fmt.Scan(&a)

	fmt.Println(conditionalConstructs_3(a))
}
