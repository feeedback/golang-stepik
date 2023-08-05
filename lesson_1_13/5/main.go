// https://stepik.org/lesson/229320/step/6
package main

import "fmt"

func isTriangleExists(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func exercises5(a, b, c int) string {
	if isTriangleExists(a, b, c) {
		return "Существует"
	}
	return "Не существует"

}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(exercises5(a, b, c))
}
