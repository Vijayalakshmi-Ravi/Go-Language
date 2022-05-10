package main

import "fmt"

func welcome() {
	fmt.Println("Welcome")
}
func wel() {
	fmt.Println("Hello")
}
func main() {
	defer welcome()
	fmt.Println("hey")
	wel()
	for i := 1; i < 5; i++ {
		defer fmt.Println(i)
	}
}
