// https://stepik.org/lesson/229320/step/2
package main

import "fmt"

// По данному трехзначному числу определите, все ли его цифры различны.
func exercises1(n int16) int16 {
	var a, b, c int16
	a = n / 100
	b = n / 10 % 10
	c = n % 10

	return a + b + c
}

func main() {
	var a int16 // натуральное трехзначное число
	fmt.Scan(&a)

	fmt.Println(exercises1(a))
}
