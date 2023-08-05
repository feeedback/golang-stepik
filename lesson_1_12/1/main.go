// https://stepik.org/lesson/228265/step/5
package main

import (
	"fmt"
)

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	for i := 0; i < 3; i++ {
		var index1, index2 int

		fmt.Scan(&index1, &index2)

		workArray[index1], workArray[index2] = workArray[index2], workArray[index1]
	}

	for _, num := range workArray {
		fmt.Print(num, " ")
	}
}
