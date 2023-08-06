package main // https://stepik.org/lesson/228838/step/8

import "fmt"

func vote(x, y, z int) int {
	result := 0
	if x+y+z >= 2 {
		result = 1
	}
	return result
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(vote(a, b, c))
}

// ***** for stepik ****
// only "vote" function
