package main

import "fmt"

func Condition(a, b int) {
	if a == b {
		disp1()
	} else if a > b {
		disp2()
	} else {
		fmt.Println("A is lesser than b")
	}
}

func disp1() {
	fmt.Println("A is equal to b")
}
func disp2() {
	fmt.Println("A is greater than b")
}
func main() {
	var status bool
	fmt.Println("Status true or false ?")
	fmt.Scanln(&status)
	var a, b int
	if status == true {
		fmt.Println("Status is set to true")
		fmt.Println("Enter values: ")
		fmt.Scanln(&a, &b)
		Condition(a, b)
		// idomatic way of go for not using else is to use return in place
		return
	}
	fmt.Println("Please set the status 'true'")
}
