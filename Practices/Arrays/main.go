package main

import "fmt"

func Aaa(num [5]int) {
	num[0] = 5
	fmt.Println("Func", num)
}
func main() {
	fmt.Println("ARRAYS")
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	b := [...]int{5, 4, 3, 2, 1}
	fmt.Println(b)
	c := b
	c[0] = 4
	fmt.Println(c)
	num := [...]int{11, 22, 33, 44, 55}
	fmt.Println("Array before func call ", num)
	Aaa(num)
	fmt.Println("Array After func call ", num, "length of array ", len(num))
	iterat()
	multi()
}

func iterat() {
	a := [...]int{10, 20, 30, 40, 50, 60, 70, 80}
	for i := 0; i < len(a); i++ {
		fmt.Printf("Index %d , Value %d\n", i, a[i])
	}
	fmt.Println(" Using range-------------------------------------")
	for i, v := range a {
		fmt.Printf("Index %d , Value %d\n", i, v)
	}
}

func multi() {
	a := [3][3]int{
		{2, 4, 6}, {3, 6, 9}, {4, 8, 16},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%d,", v2)
		}
		fmt.Printf("\n")
	}
	res := 0
	nums := [3]int{2, 4, 6}
	for i, v := range nums {
		if i%2 == 0 {
			res += v
		}
	}
	fmt.Println(res)
}
