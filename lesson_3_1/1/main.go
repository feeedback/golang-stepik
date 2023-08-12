package main // https://stepik.org/lesson/345543/step/5

import "fmt"

func work(x int) int {
	return x * 2
}

func main() {
	// for stepik only code inside main fn
	var numbers [10]int

	for i := 0; i < 10; i++ {
		fmt.Scan(&numbers[i])
	}

	cacheMap := make(map[int]int)

	for _, num := range numbers {
		if _, exists := cacheMap[num]; !exists {
			cacheMap[num] = work(num)
		}
		fmt.Printf("%d ", cacheMap[num])
	}
}
