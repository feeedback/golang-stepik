package main // https://stepik.org/lesson/228261/step/16

import "fmt"

func arithmeticCalculations_6(degree int) string {
	const degreeByMinute = 2 // (12 * 60) / 360
	const degreeByHours = 60 / degreeByMinute

	return fmt.Sprintf("It is %v hours %v minutes.", degree/degreeByHours, (degree%degreeByHours)*degreeByMinute)
}

func main() {
	var a int // (0 < d < 360)
	fmt.Scan(&a)

	fmt.Println(arithmeticCalculations_6(a))
}
