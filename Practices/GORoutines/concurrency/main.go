package main

import (
	"fmt"
	//"time"
)

func display(a string) {
	for i := 0; i < 3; i++ {
		//time.Sleep(1 * time.Second)
		fmt.Println(a)
	}
}

func main() {

	go display("Welcome")

	display("Viji")
}
