package main // https://stepik.org/lesson/232593/step/6

import "fmt"

// По данному трехзначному числу определите, все ли его цифры различны.
func conditionalConstructs_2(n int) string {
	var a, b, c int
	a = n / 100
	b = n / 10 % 10
	c = n % 10

	if a != b && b != c && a != c {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var a int // натуральное трехзначное число
	fmt.Scan(&a)

	fmt.Println(conditionalConstructs_2(a))
}
