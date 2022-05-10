package main

import "fmt"

func find(a int, nums ...int) {
	found := false
	for i, v := range nums {
		if v == a {
			fmt.Printf("The number %d found at index %d", v, i)
			found = true
		}
	}
	if !found {
		fmt.Println("The number not found")
	}

}

func change(a ...string) {
	a[0] = "Hello"
	a = append(a, "Bye")
	fmt.Println(a)
}

func main() {
	nums := []int{10, 40, 20, 30, 50}
	find(50, nums...)
	s := []string{"Hi", "Welcome"}
	change(s...)
	fmt.Println(s)
}
