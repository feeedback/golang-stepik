package main // https://stepik.org/lesson/345547/step/13

import (
	"fmt"
	"time"
)

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	calculatorChan := calculator(firstChan, secondChan, stopChan)

	go func() {
		firstChan <- 2
		secondChan <- 4
	}()

	fmt.Println("First result:", <-calculatorChan)
	fmt.Println("Second result:", <-calculatorChan)

	close(stopChan)

	time.Sleep(time.Second)
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outCh := make(chan int)

	go func() {
		defer close(outCh)

		for {
			select {
			case val := <-firstChan:
				outCh <- val * val

			case val := <-secondChan:
				outCh <- val * 3

			case <-stopChan:
				return
			default:
				return
			}
		}
	}()

	return outCh
}
