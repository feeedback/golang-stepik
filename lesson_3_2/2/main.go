package main // https://stepik.org/lesson/348563/step/13
import (
	"strconv"
	"strings"
	"unicode"
)

func filterOnlyDigit(r rune) bool { return !unicode.IsDigit(r) }

func adding(aRaw, bRaw string) int64 {
	a := strings.TrimFunc(aRaw, filterOnlyDigit)
	b := strings.TrimFunc(bRaw, filterOnlyDigit)

	numA, _ := strconv.Atoi(a)
	numB, _ := strconv.Atoi(b)

	sum := int64(numA + numB)
	return sum
}
