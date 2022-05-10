package main

import "fmt"

func Odd(i int) {
	if i <= 5 {
		fmt.Println("Number sequence 1-5 ")
		fmt.Println("ODD no: ", i)
	} else {
		fmt.Println("Number sequence 6-10")
		fmt.Println("ODD no: ", i)
	}
}

func Even(i int) {
	if i <= 5 {
		fmt.Println("Number sequence 1-5 ")
		fmt.Println("EVEN no: ", i)
	} else {
		fmt.Println("Number sequence 6-10")
		fmt.Println("EVEN no: ", i)
	}
}
func main() {
	//infinite loop
	//--------------
	// for {
	// 	fmt.Println("HEYYYYYYYYYY")
	// }

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			Odd(i)
			continue
		}
		Even(i)
	}
	fmt.Println("Numbers :")
outer:
	for i := 0; i <= 5; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i=%d,j=%d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
}
