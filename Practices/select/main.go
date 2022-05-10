package main

import (
	"fmt"
	"time"
)

func addSum(a, b int, ch chan int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	ch <- sum
}
func mulSum(a, b int, ch chan int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i * i
	}
	ch <- sum
}
func main() {
	//channel initilization
	addch := make(chan int)
	mulch := make(chan int)

	//go routines
	go addSum(1, 50, addch)
	go mulSum(1, 10, mulch)

	//celect channel
	for {
		select {
		case x := <-addch:
			fmt.Println(x)
			return
		case y := <-mulch:
			fmt.Println(y)
			return
		default:
			fmt.Println("*data coming nthinfg avail now")
			time.Sleep(50 * time.Millisecond)
		}

	}
}
