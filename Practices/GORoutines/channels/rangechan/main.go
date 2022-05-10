package main

import (
	"fmt"
	"time"
)

var s1 *int

func myfunc(ch chan *int) {
	fmt.Println(*<-ch)
}

func main() {

	a, b := 5, 10
	var s *int = &a
	ch := make(chan *int)
	go myfunc(ch)
	ch <- s
	s = &b
	go myfunc(ch)
	ch <- s
	time.Sleep(1 * time.Second)

}
