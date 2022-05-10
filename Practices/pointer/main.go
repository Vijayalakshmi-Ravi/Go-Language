package main

import "fmt"

func change(val *int) *int {
	*val = 55
	return val
}
func modify(arr []int) {
	arr[0] = 90
}
func main() {
	b := 401
	var a *int = &b
	fmt.Printf("Type of a %T\n", a)
	fmt.Println("Address of b", a)
	size := new(int)
	fmt.Printf("Value %d type %T address %v", *size, size, size)
	*size = 85
	fmt.Println("new size ", *size)
	v := 52.20
	var fPointer *float64 = &v
	fmt.Printf("Type %T Value %5f\n", fPointer, v)
	fmt.Println(fPointer)
	ret := change(a)
	fmt.Println("value of b after passing function", *ret)

	//--------------------------------------------------------------------
	arr := [3]int{40, 60, 80}
	modify(arr[:])
	fmt.Println(arr)
}
