package main

import (
	"fmt"
	"time"
)

func calc(ch chan int) {

	fmt.Println(100 + <-ch)
}
func main() {

	ch1 := make(chan int)

	ch2 := make(chan int)

	go calc(ch1)

	ch1 <- 3

	go calc(ch2)

	ch2 <- 5

	go calc(ch1)
	ch1 <- 10

	time.Sleep(1 * time.Second)

}
