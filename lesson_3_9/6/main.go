package main

import (
	"fmt"
	"math/rand" // only for test
	"sync"
	"time"
)

type Worker struct {
	sync.Mutex
	sync.WaitGroup
}

func (w *Worker) Do(workFn func(int) int, argNum int, sum *int) {
	result := workFn(argNum)

	w.Lock()
	*sum += result
	w.Unlock()

	w.Done()
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		var w Worker

		results := make([]int, n)

		for i := range results {
			w.Add(2)

			go w.Do(fn, <-in1, &results[i])
			go w.Do(fn, <-in2, &results[i])
		}

		w.Wait()

		for _, sum := range results {
			out <- sum
		}
	}()
}

// only for test
const N = 20

func main() {
	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Millisecond)

		return x * 2
	}

	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()

	go merge2Channels(fn, in1, in2, out, N+1)

	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	evenFail := false

	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			evenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}

	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if evenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)

	if duration.Milliseconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}
