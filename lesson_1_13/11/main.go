package main // https://stepik.org/lesson/229320/step/12

import "fmt"

func calculatePlural(n int) string {
	plural := "many"

	n1 := n % 10
	n2 := n % 100

	if n1 == 1 && n2 != 11 {
		plural = "one"
	} else if n1 >= 2 && n1 <= 4 && (n2 < 10 || n2 >= 20) {
		plural = "few"
	}
	return plural
}

func exercises11(n int) string {
	plural := calculatePlural(n)
	word := "korov"

	switch plural {
	case "one":
		word = "korova"
	case "few":
		word = "korovy"
	case "many":
		word = "korov"
	}

	return fmt.Sprintf("%d %s", n, word)
}

func main() {
	var a int // (0<n<100).
	fmt.Scan(&a)

	fmt.Println(exercises11(a))
}
