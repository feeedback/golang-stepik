package main // https://stepik.org/lesson/228263/step/8

import (
	"fmt"
)

func cycles_6(nums []int) []int {
	var numsOutput []int

	for i := 0; i < len(nums) && nums[i] <= 100; i++ {
		if nums[i] >= 10 {
			numsOutput = append(numsOutput, nums[i])
		}
	}

	return numsOutput
}

func main() {
	var n int
	var nums []int

	for fmt.Scanln(&n); n <= 100; fmt.Scanln(&n) {
		nums = append(nums, n)
	}

	for _, v := range cycles_6(nums) {
		fmt.Println(v)
	}
}
