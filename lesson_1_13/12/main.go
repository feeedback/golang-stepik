package main // https://stepik.org/lesson/229320/step/13

import (
	"fmt"
	"golang-stepik/utils"
	"strings"
)

func exercises12(n int) []int {
	var nums []int

	for i := 1; i <= n; i *= 2 {
		nums = append(nums, i)
	}

	return nums
}

func main() {
	var a int
	fmt.Scan(&a)

	output := exercises12(a)

	mappedNums := utils.MapFn(output, utils.MapFnType[int, string](func(n int) string {
		return fmt.Sprintf("%d", n)
	}))
	fmt.Println(strings.Join(mappedNums, " "))

	// // version without helper fn

	// mappedNums := make([]string, len(output))
	// for i, num := range output {
	// 	mappedNums[i] = fmt.Sprintf("%d", num)
	// }
}
