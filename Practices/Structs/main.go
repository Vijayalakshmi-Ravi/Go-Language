package main

import (
	"STRUCTS/emp"
	"fmt"
)

type ContactDetails struct {
	PhNos  int
	MailID string
	emp.Employee
}

func main() {
	emp1 := emp.Employee{
		ID:   1001,
		Age:  24,
		Name: "Vj",
	}
	fmt.Println("ID: ", emp1.ID)
	fmt.Println("NAME: ", emp1.Name)
	fmt.Println("AGE: ", emp1.Age)
	c := ContactDetails{
		PhNos:  7894561230,
		MailID: "sdfgh@gmail.com",
		Employee: emp.Employee{
			ID:   emp1.ID,
			Name: emp1.Name,
			Age:  emp1.Age,
		},
	}
	fmt.Println("EMPLOYEE DETAILS\n*************************\n", "\nID: ", c.ID, "\nName: ", c.Name, "\nAge: ", c.Age, "\nPh: ", c.PhNos, "\nEmailID: ", c.MailID)
	c2 := &ContactDetails{
		PhNos:  9874563210,
		MailID: "qwerty@gmail.com",
		Employee: emp.Employee{
			ID:   1002,
			Name: "DV",
			Age:  24,
		},
	}
	c2.Age = 25
	fmt.Println("EMPLOYEE DETAILS\n*************************\n", "\nID: ", c2.ID, "\nName: ", c2.Name, "\nAge: ", c2.Age, "\nPh: ", c2.PhNos, "\nEmailID: ", c2.MailID)
}
