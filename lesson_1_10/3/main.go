package main // https://stepik.org/lesson/228263/step/4

import "fmt"

// находит сумму двузначных чисел, кратных 8
func cycles_3(n int, nums []int) int {
	sum := 0

	for i := 0; i < n; i++ {
		num := nums[i]

		if 10 <= num && num <= 99 && num%8 == 0 {
			sum += num
		}
	}
	return sum
}

func main() {
	var n int
	var nums []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var word int
		fmt.Scan(&word)

		nums = append(nums, word)
	}

	fmt.Println(cycles_3(n, nums))
}
