package main // https://stepik.org/lesson/360357/step/8

func main() {
	ch := make(chan int)

	go task(ch, 5)

	result := <-ch

	println(result)
}

// for stepik only below part
func task(ch chan int, N int) {
	ch <- N + 1
}
