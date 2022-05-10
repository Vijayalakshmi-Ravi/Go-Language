package main

import (
	"fmt"
)

func calc1(a, b int) int {
	var t1 = a + b
	return t1
}

func calc2(c int, d float32) float32 {
	var t2 float32 = float32(c) + d
	return t2
}

func calc3(l, b int) (int, int) {
	var area = l * b
	var peri = (l * b) * 2
	return area, peri
}

func calc4(e int, f float32) (area, peri float32) {
	area = float32(e) * f
	peri = (float32(e) * f) * 2
	return
}

func main() {
	len, bred := 5, 10

	fmt.Println("Total: ", calc1(5, 10))
	fmt.Println("Total: ", calc2(5, 5.20))
	fmt.Printf("Area : % d & Peri : %d", len, bred)
	fmt.Println(calc4(5, 5))
	ar, _ := calc4(5, 5)
	fmt.Println("Only Area", ar)
}
