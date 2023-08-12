package main // https://stepik.org/lesson/264473/step/4

import "fmt"

func divide(a, b int) (int, error) {
	return 0, nil
}

// for stepik only main fn
func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}

	result, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}

	fmt.Println(result)
}
