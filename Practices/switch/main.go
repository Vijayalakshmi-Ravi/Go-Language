package main

import (
	"fmt"
	"math/rand"
)

func Num(n int) int {
	n *= 5
	return n
}
func main() {
	var a int
	fmt.Println("Enter NUM: ")
	fmt.Scanln(&a)
	switch Num(a); {
	case a < 10:
		fmt.Printf("The integer is %d \n", a)
	case a < 20:
		fmt.Printf("The integer is %d \n", a)
	case a < 30:
		fmt.Printf("The integer is %d \n", a)
	case a < 40:
		fmt.Printf("The integer is %d \n", a)
	case a <= 50:
		fmt.Printf("The integer is %d \n", a)
	default:
		fmt.Println("Number is less than 10 or greater than 50")
	}

	switch n := 10; {
	case n < 50:
		fmt.Printf("Num %d is less than 50\n", n)
		fallthrough

	case n < 100:
		fmt.Printf("Num %d is less than 100\n", n)
		//fallthrough---> if fall through is given begore default then default also will execute

	default:
		fmt.Println("Num is greater than 100")
	}

randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 != 0:
			fmt.Printf("ODD NUM %d \n", i)
			//breaks out of only switch
			break
		case i%2 == 0:
			fmt.Printf("EVEN NUm %d \n", i)
			//breaks out of even for loop
			break randloop
		}
	}
}
