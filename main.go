package main

import "fmt"

//	"time"

/*
func main() {
	in := make(chan int)

	go func() {
		var i int
		for {
			a := squareTheNumer(i)
			in <- a
			i++
		}
	}()
	for res := range in {
		fmt.Println("result is :- ", res)
	}
}
func squareTheNumer(i int) int {
	time.Sleep(1 * time.Second)
	return i * i
}
*/
func main() {
	in := make(chan int)
	done := make(chan bool)
	go producer(in, done)
	go consumer(in)
	<-done
}

func producer(in chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		in <- i * i
	}
	done <- true
}
func consumer(in chan int) {
	for res := range in {
		fmt.Println("result is :- ", res)
	}

}
