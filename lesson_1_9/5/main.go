package main // https://stepik.org/lesson/232593/step/9

import "fmt"

func conditionalConstructs_5(n int) string {
	if n%400 == 0 || (n%100 != 0 && n%4 == 0) {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var a int // номер года (целое, положительное, не превышает 10000)
	fmt.Scan(&a)

	fmt.Println(conditionalConstructs_5(a))
}
