// https://stepik.org/lesson/228265/step/14
package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	max := array[0]

	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	fmt.Println(max)
}
