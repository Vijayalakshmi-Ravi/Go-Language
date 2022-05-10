package main

import (
	"fmt"
	"time"
)

func out(from, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
	ch <- true
	close(ch)
}
func evenSum(from, to int, ch chan int) {
	res := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			res += 1
		}
	}
	ch <- res
}
func squareSum(from, to int, ch chan int) {
	res := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			res += i * i
		}
	}
	ch <- res
}
func main() {
	ch := make(chan bool)
	go out(0, 5, ch)
	go out(6, 10, ch)
	<-ch
	ch1 := make(chan int)
	ch2 := make(chan int)
	go evenSum(0, 100, ch1)
	go squareSum(0, 100, ch2)
	fmt.Println(<-ch1 + <-ch2)

}
