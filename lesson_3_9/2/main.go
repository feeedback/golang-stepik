package main // https://stepik.org/lesson/345547/step/5

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			work()

			defer wg.Done()
		}()
	}

	wg.Wait()
}

func work() {
	fmt.Printf("Горутина начала выполнение \n")
	time.Sleep(2 * time.Second)
	fmt.Printf("Горутина завершила выполнение \n")
}
