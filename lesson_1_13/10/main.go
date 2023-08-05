package main // https://stepik.org/lesson/229320/step/11

import "fmt"

func exercises10(a, b int) int {
	maxMultipleOf7 := -1

	for num := b; num >= a; num -= 1 {
		if num%7 == 0 {
			maxMultipleOf7 = num
			break
		}
	}
	return maxMultipleOf7
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	output := exercises10(a, b)

	if output == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println(output)
	}
}
