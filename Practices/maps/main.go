package main

import "fmt"

type Employee struct {
	state   string
	country string
}

func main() {
	EmployeeSalary := map[string]int{
		"Vj":  28250,
		"Dv":  28250,
		"Jas": 45082,
	}

	fmt.Println(EmployeeSalary)
	Employee1 := "Vj"
	Sal1 := EmployeeSalary[Employee1]
	fmt.Println("Salary of Employee", Employee1, " is ", Sal1)
	value, ok := EmployeeSalary[Employee1]
	if ok == true {
		fmt.Println("Salary of ", Employee1, "is ", value)
	} else {
		fmt.Println("not found")
	}

	for k, v := range EmployeeSalary {
		fmt.Println(k, ":", v)
	}

	delete(EmployeeSalary, "Dv")

	for k, v := range EmployeeSalary {
		fmt.Println(k, ":", v)
	}

	emp1 := Employee{
		state:   "Washington D.C",
		country: "USA",
	}
	emp2 := Employee{
		state:   " New Jersey",
		country: "USA",
	}

	EmployeeInfo := map[string]Employee{
		"Vj": emp1,
		"Dv": emp2,
	}
	for name, info := range EmployeeInfo {
		fmt.Printf("Employee Name %s State %s country %s\n", name, info.state, info.country)
	}
	fmt.Println("length of EmployeeInfo", len(EmployeeInfo))
	EmployeeSalary["vj"] = 58000
	fmt.Println(EmployeeSalary)
}
