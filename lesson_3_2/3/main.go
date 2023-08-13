package main // https://stepik.org/lesson/348563/step/14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseNumber(s string) (float64, error) {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", ".")

	return strconv.ParseFloat(s, 64)
}

func divide(csvLine string) string {
	numsStr := strings.Split(csvLine, ";")

	num1, _ := parseNumber(numsStr[0])
	num2, _ := parseNumber(numsStr[1])

	roundedQuotient := fmt.Sprintf("%.4f", float64(num1)/float64(num2))
	return roundedQuotient
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := scanner.Text()

	fmt.Println(divide(a))
}
