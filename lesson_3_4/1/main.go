package main // https://stepik.org/lesson/350788/step/10

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func readTask() (value1, value2, operation interface{}) {
	return 10.5, 5.2, "/"
}

// below for stepik
func performMathOperation(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation")
	}
}

func checkTypes(value1, value2, operation interface{}) (float64, float64, string, error) {
	a, aOk := value1.(float64)
	b, bOk := value2.(float64)

	if !aOk {
		return 0, 0, "", fmt.Errorf(fmt.Sprintf("value=%v: %T", value1, value1))
	}
	if !bOk {
		return 0, 0, "", fmt.Errorf(fmt.Sprintf("value=%v: %T", value2, value2))
	}

	opMark, operationOk := operation.(string)
	if !operationOk {
		return 0, 0, "", fmt.Errorf("неизвестная операция")
	}
	return a, b, opMark, nil
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	a, b, opMark, err := checkTypes(value1, value2, operation)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resNum, err := performMathOperation(a, b, opMark)
	if err != nil {
		fmt.Println("неизвестная операция")
		return
	}

	fmt.Printf("%.4f", resNum)
}
