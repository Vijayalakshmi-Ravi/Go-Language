package area

import "fmt"

func CalcSQ(l int) (total int) {
	fmt.Println("The area of Square")
	return l * l
}

func CalcRECT(l, b int) (total int) {
	fmt.Println("The area of Rectangle")
	return l * b
}
