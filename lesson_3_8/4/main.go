package main // https://stepik.org/lesson/360357/step/10

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string, 10)

	go removeDuplicates(inputStream, outputStream)

	for _, value := range []string{"1", "2", "2", "1", "1", "3", "3", "3"} {
		inputStream <- value
	}
	close(inputStream)

	for value := range outputStream {
		println(value)
	} // 1 2 1 3
}

func removeDuplicatesAll(inputCh <-chan string, outputCh chan<- string) {
	defer close(outputCh)

	set := make(map[string]bool)

	for val := range inputCh {
		if !set[val] {
			set[val] = true

			outputCh <- val
		}
	}
}

// for stepik only below part
func removeDuplicates(inputCh <-chan string, outputCh chan<- string) {
	defer close(outputCh)

	var lastValue string

	for val := range inputCh {
		if val != lastValue {
			lastValue = val

			outputCh <- val
		}
	}
}
