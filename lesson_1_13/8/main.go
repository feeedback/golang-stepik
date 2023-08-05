package main // https://stepik.org/lesson/229320/step/9

import (
	"fmt"
	"math"
)

func countMinElements(nums []int) int {
	min := math.MaxInt64
	countMaxEl := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			countMaxEl = 1
		} else if nums[i] == min {
			countMaxEl += 1
		}
	}
	return countMaxEl
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(countMinElements(nums))
}
