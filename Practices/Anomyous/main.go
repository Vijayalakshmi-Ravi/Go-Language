package main

import "fmt"

type Employee struct {
	eid   int
	ename string
	dept  string
}

func main() {

	// Anonymous function
	func() {

		fmt.Println("Welcome! to GeeksforGeeks")
	}()
	fmt.Println("huhdius")

	func() {

		fmt.Println("Welcome! to GeeksforGeeks")
	}()

	var arr [2]Employee
	arr[0] = Employee{
		eid:   1,
		ename: "VJ",
	}
	arr[1] = Employee{
		eid:   2,
		ename: "DV",
	}

	fmt.Println(arr[1])
}
