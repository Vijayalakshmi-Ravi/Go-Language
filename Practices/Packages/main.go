package main

import (
	"PACKAGES/area"
	"PACKAGES/mathfunc"
	"fmt"
)

var l, b, n int

func main() {	
	fmt.Println("Enter num to find Area ")
	fmt.Scanln(&l, &b)
	fmt.Println("AREA OF SQUARE", area.CalcSQ(l))
	fmt.Println("AREA OF RECTANGLE", area.CalcRECT(l, b))
	fmt.Println("Enter num to find SQRt ")
	fmt.Scanln(&n)
	fmt.Println(mathfunc.Sqroot(n))
}
