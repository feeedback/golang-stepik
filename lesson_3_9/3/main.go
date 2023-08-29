package main // https://stepik.org/lesson/345547/step/6

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var num int
	var mu sync.Mutex // Мьютекс для синхронизации доступа к num
	wg := new(sync.WaitGroup)

	startTime := time.Now()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			time.Sleep(time.Millisecond * 10) // задержка вне Lock

			mu.Lock() // Захватываем мьютекс перед изменением num
			num++
			mu.Unlock() // Освобождаем мьютекс после изменения num
		}(wg)
	}

	wg.Wait()

	fmt.Println(num)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	elapsedNanoseconds := elapsedTime.Seconds()

	fmt.Println("Time elapsed:", elapsedNanoseconds, "sec")
}
