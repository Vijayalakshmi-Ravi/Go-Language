package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Student struct {
	id      string
	name    string
	mark1   string
	mark2   string
	mark3   string
	mark4   string
	mark5   string
	total   string
	average string
	grade   string
}

var flag = true
var rows int
var student [10][1]Student
var option string
var ques [7]string = [7]string{"ID ", "Name ", "Subject 1 ", "Subject 2 ", "Subject 3 ", "Subject 4 ", "Subject 5 "}

func addStudent() {
	for i := 0; i < 10; i++ {
		for j := 0; j < len(ques); j++ {
			fmt.Println("Enter ", ques[j])
			if j == 0 {
				fmt.Scanln(&student[i][0].id)
				sid, err := strconv.Atoi(student[i][0].id)
				if err != nil {
					fmt.Println("ID Must be an Integer")
					j--
				} else if sid == 0 {
					fmt.Println("ID cannot be ZERO")
					j--
				} else if sid < 0 {
					fmt.Println("Studnet Id cannot be Negative")
					j--
				}
			}
			if j == 1 {
				fmt.Scanln(&student[i][0].name)
				re := regexp.MustCompile("[a-zA-Z]$")
				if !re.MatchString(student[i][0].name) {
					fmt.Println("Name can only be Alphabetic")
					j--
				}
			}
			if j == 2 {
				fmt.Scanln(&student[i][0].mark1)
				y, err := strconv.Atoi(student[i][0].mark1)
				if err == nil {
					if y > 100 {
						fmt.Println("Marks cannot be greater than 100")
						j--
					} else if y < 0 {
						fmt.Println("Marks cannot be Negative")
						j--
					} else {
						student[i][0].total += strconv.Itoa(y)
					}
				} else {
					fmt.Println("Marks accepts only Integer")
					j--
				}
			}
			if j == 3 {
				fmt.Scanln(&student[i][0].mark2)
				y, err := strconv.Atoi(student[i][0].mark2)
				if err == nil {
					if y > 100 {
						fmt.Println("Marks cannot be greater than 100")
						j--
					} else if y < 0 {
						fmt.Println("Marks cannot be Negative")
						j--
					} else {
						student[i][0].total += strconv.Itoa(y)
					}
				} else {
					fmt.Println("Marks accepts only Integer")
					j--
				}
			}
			if j == 4 {
				fmt.Scanln(&student[i][0].mark3)
				y, err := strconv.Atoi(student[i][0].mark3)
				if err == nil {
					if y > 100 {
						fmt.Println("Marks cannot be greater than 100")
						j--
					} else if y < 0 {
						fmt.Println("Marks cannot be Negative")
						j--
					} else {
						student[i][0].total += strconv.Itoa(y)
					}
				} else {
					fmt.Println("Marks accepts only Integer")
					j--
				}
			}
			if j == 5 {
				fmt.Scanln(&student[i][0].mark4)
				y, err := strconv.Atoi(student[i][0].mark4)
				if err == nil {
					if y > 100 {
						fmt.Println("Marks cannot be greater than 100")
						j--
					} else if y < 0 {
						fmt.Println("Marks cannot be Negative")
						j--
					} else {
						student[i][0].total += strconv.Itoa(y)
					}
				} else {
					fmt.Println("Marks accepts only Integer")
					j--
				}
			}
			if j == 6 {
				fmt.Scanln(&student[i][0].mark5)
				y, err := strconv.Atoi(student[i][0].mark5)
				if err == nil {
					if y > 100 {
						fmt.Println("Marks cannot be greater than 100")
						j--
					} else if y < 0 {
						fmt.Println("Marks cannot be Negative")
						j--
					} else {
						student[i][0].total += strconv.Itoa(y)
					}
				} else {
					fmt.Println("Marks accepts only Integer")
					j--
				}
			}
		}
	}
}

func viewAll() {
	fmt.Println("view")
	fmt.Println(student)
}

func searchStudent() {
	fmt.Println("search")
}
func subMenu() {
	fmt.Println("Select any one \n1. Main Menu 2. Press any key to exit")
	fmt.Scanln(&option)
	noption, _ := strconv.Atoi(option)
	if noption == 1 {
		menu()
	} else {
		flag = false
		fmt.Println("Thank you!")
	}
}
func menu() {
	for flag {
		fmt.Println("Select One option \n1. Add student 2. View all 3. Search ")
		fmt.Scanln(&option)
		noption, _ := strconv.Atoi(option)
		switch noption {
		case 1:
			addStudent()
		case 2:
			viewAll()
		case 3:
			searchStudent()
		default:
			fmt.Println("Invalid choice\n ")
			subMenu()
		}
	}
}

func main() {

	for i, v := range student {
		fmt.Println(i, "-", v)
	}
	fmt.Println("Welcome \n ")
	menu()

}
