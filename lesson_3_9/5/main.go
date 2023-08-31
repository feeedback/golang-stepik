package main // https://stepik.org/lesson/345547/step/14

func main() {
	arguments := make(chan int)
	done := make(chan struct{})

	resCh := calculator(arguments, done)

	go func() {
		for i := 1; i <= 10; i++ {
			arguments <- i
		}

		close(arguments)
		done <- struct{}{}
	}()

	result := <-resCh
	println("Sum:", result)
}

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	outChan := make(chan int)

	go func() {
		defer close(outChan)
		sum := 0

		for {
			select {
			case arg := <-arguments:
				sum += arg

			case <-done:
				outChan <- sum
				return
			}
		}
	}()

	return outChan
}
