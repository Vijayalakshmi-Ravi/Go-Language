package main

import "fmt"

func myfunc(ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- "Good Morning"
	}
	close(ch)
}

func main() {

	ch := make(chan string)

	go myfunc(ch)

	for {
		res, ok := <-ch
		if ok == false {
			fmt.Println("Channel is closed", ok)
			break
		}
		fmt.Println("channel open", res, ok)
	}

}
