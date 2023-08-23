package main // https://stepik.org/lesson/351892/step/9

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		sum += num
	}

	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
