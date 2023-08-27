package main // https://stepik.org/lesson/360357/step/9

func main() {
	ch := make(chan string)

	go task2(ch, "")

	result := <-ch

	println(result)
}

// for stepik only below part
func task2(ch chan string, str string) {
	ch <- str + " "
	ch <- str + " "
	ch <- str + " "
	ch <- str + " "
	ch <- str + " "
}
