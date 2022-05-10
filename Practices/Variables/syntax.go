package main

import "fmt"

func main() {
	var a int
	fmt.Println("a=", a)

	var b = 5
	fmt.Println("b=", b)

	var c int
	fmt.Println("c=", c)
	c = 5
	fmt.Println("c=", c)
	c = 40
	fmt.Println("c=", c)

	var h, e int = 10, 55 //var h,e=10,55 also applicable
	fmt.Println(h, e)

	var (
		age    = 10
		height = 152
	)
	fmt.Println(age, height)

	count := 10
	fmt.Println("count:", count)

	cout, jac := 10, 50
	cout, iuy := 15, 20

	fmt.Println("cout,jac", cout, jac)
	fmt.Println("cout,iuy", cout, iuy)
}
