package main // https://stepik.org/lesson/345547/step/3

import (
	"fmt"
	"time"
)

func work() {
	fmt.Println("start...")

	time.Sleep(time.Second * 2)

	fmt.Println("finish...")
}

func main() {
	done := make(chan struct{})

	go func() {
		work()

		close(done)
	}()
	<-done
}

// func main() {
// 	workDone := false

// 	go func() {
// 		work()
// 		workDone = true
// 	}()

// 	for !workDone {
// 		time.Sleep(time.Millisecond * 100)
// 	}
// }
