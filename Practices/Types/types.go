package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := false
	c := a && b
	fmt.Println(c)
	bi := 3
	fmt.Printf("type of bi is %T, size of bi %d", bi, unsafe.Sizeof(bi))

	i := 5.25
	j := 75.25
	fmt.Printf("\ntype of i is %T,type of j is %T", i, j)

	com1 := complex(4, 5)
	com2 := 5 + 2i
	fmt.Println("sum ", com1+com2)
	fmt.Println("Product", com1*com2)
}
