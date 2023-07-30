package main // https://stepik.org/lesson/228263/step/5

import (
	"fmt"
	"math"
)

func countMaxElements(nums []int) int {
	max := math.MinInt64
	countMaxEl := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			countMaxEl = 1
		} else if nums[i] == max {
			countMaxEl += 1
		}
	}

	return countMaxEl
}

func main() {
	var n int
	var nums []int

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		nums = append(nums, n)
	}

	fmt.Println(countMaxElements(nums))
}
