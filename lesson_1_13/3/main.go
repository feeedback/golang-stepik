// https://stepik.org/lesson/229320/step/4
package main

import "fmt"

func exercises3(k int) string {
	hours := k / 3600
	minutes := (k % 3600) / 60

	return fmt.Sprintf("It is %d hours %d minutes.", hours, minutes)
}

func main() {
	var a int // положительное целое число
	fmt.Scan(&a)

	fmt.Println(exercises3(a))
}
