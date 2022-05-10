package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welocme")

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("Viji")
		}
	}()

	time.Sleep(9 * time.Second)
	fmt.Println("Bye")
}
