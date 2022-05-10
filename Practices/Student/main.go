package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var rows int
var student [10][12]string
var ques [7]string = [7]string{"Roll Number ", "Name ", "Subject 1 ", "Subject 2 ", "Subject 3 ", "Subject 4 ", "Subject 5 "}
var choice string
var flag = true

func menu() {

	for flag {
		choice = ""
		fmt.Println("\nEnter Choice \n1.Add Student \n2.View All \n3.Search Student")
		fmt.Scanln(&choice)
		fmt.Println()
		schoice, _ := strconv.Atoi(choice)
		fmt.Println()
		if schoice == 1 {
			addStudent()
		} else if schoice == 2 {
			viewAll()
		} else if schoice == 3 {
			searchStudent()
		} else {
			fmt.Println("You have entered Invaild Option", choice)
			subMenu()
		}
	}
}

func subMenu() {
	fmt.Println("Select any one \n1.Main Menu \n2.Presss any key to exit")
	fmt.Scanln(&choice)
	schoice, _ := strconv.Atoi(choice)
	if schoice == 1 {
		menu()
	} else {
		flag = false
		fmt.Println("Thank you!")
	}
}

func addStudent() {
	for i := 0; i < len(ques); i++ {
		fmt.Println("Enter" + ques[i])
		fmt.Scanln(&student[rows][i])
		if i == 0 {
			roll, err := strconv.Atoi(student[rows][0])
			if err != nil {
				fmt.Println("Student Id Must be a Integer")
				i--
			} else if roll == 0 {
				fmt.Println("Studnet Id cannot be Zero")
				i--
			} else if roll < 0 {
				fmt.Println("Studnet Id cannot be Negative")
				i--
			}
		outer:
			for r := 0; r < 10 && r != rows; r++ {
				for c := 0; c < 12; c++ {
					if student[r][c] == student[rows][0] {
						fmt.Println("Student ID already Exists.")
						i--
						break outer
					}
				}
			}
		}
		if i == 1 {
			// n, err := strconv.Atoi(student[rows][1])
			// if err == nil && n != 0 {
			// 	fmt.Println("Name can only be alphabetic")
			// 	i--
			// }
			re := regexp.MustCompile("^[a-zA-Z]$")
			if !re.MatchString(student[rows][1]) {
				fmt.Println("Name can only be Alphabetic")
				i--
			}
		}
		if i >= 2 {
			y, err := strconv.Atoi(student[rows][i])
			if err == nil {
				if y > 100 {
					fmt.Println("Marks cannot be greater than 100")
					i--
				} else if y < 0 {
					fmt.Println("Marks cannot be lesser than 0")
					i--
				} else {
					if y < 50 {
						y10, _ := strconv.Atoi(student[rows][10])
						y10 += 1
						student[rows][10] = strconv.Itoa(y10)
						student[rows][11] = student[rows][11] + "," + ques[i]
					}
					y1, _ := strconv.Atoi(student[rows][7])
					student[rows][7] = strconv.Itoa(y1 + y)
				}
			} else {
				fmt.Println("Marks accepts only Integer")
				i--
			}
		}
	}
	y2, _ := strconv.Atoi(student[rows][7])
	student[rows][8] = strconv.Itoa(y2 / 5)
	y3, _ := strconv.Atoi(student[rows][8])
	if y3 > 90 {
		student[rows][9] = "O Grade"
	} else if y3 > 80 {
		student[rows][9] = "A Grade"
	} else if y3 > 70 {
		student[rows][9] = "B Grade"
	} else if y3 > 60 {
		student[rows][9] = "C Grade"
	} else if y3 >= 50 {
		student[rows][9] = "D Grade"
	} else {
		student[rows][9] = "E Grade"
	}
	rows++
	fmt.Println("1 Record Added Successfully")
}

func viewAll() {
	if student[0][0] != "" {
		fmt.Println("\nStudent Details\n____________________")
		for i := 0; i < len(student); i++ {
			fmt.Println()
			for j := 0; j <= 11; j++ {
				if student[i][j] != "" {
					if j < 7 {
						fmt.Print(ques[j], ": ", student[i][j], " |")
					} else if j == 7 {
						fmt.Print("Total: ", student[i][j], " |")
					} else if j == 8 {
						fmt.Print("Average: ", student[i][j], "| ")
					} else if j == 9 {
						fmt.Print("Grade: ", student[i][j], " |")
					} else if j == 10 {
						fmt.Print("No of failed subjects: ", student[i][j], "| ")
					} else {
						fmt.Print("Failed Subject list: ", student[i][j])
					}
				}

			}
		}
	} else {
		fmt.Println("\nNo Records, Database is empty")
	}
}

func searchStudent() {
	if student[0][0] != "" {
		fmt.Println("Enter STUDENT ID to be searched:")
		var searchID int
		fmt.Scanln(&searchID)
		fmt.Println("\nStudent Record: \n____________________")
		for i := 0; i < len(student); i++ {
			sid, _ := strconv.Atoi(student[i][0])
			if sid == searchID {
				fmt.Println()
				for j := 0; j <= 11; j++ {
					if j < 7 {
						fmt.Print(ques[j], ": ", student[i][j], " |")
					} else if j == 7 {
						fmt.Print("Total: ", student[i][j], " |")
					} else if j == 8 {
						fmt.Print("Average: ", student[i][j], "| ")
					} else if j == 9 {
						fmt.Print("Grade: ", student[i][j], " |")
					} else if j == 10 {
						fmt.Print("No of failed subjects: ", student[i][j], "| ")
					} else {
						fmt.Print("Failed Subject list: ", student[i][j])
					}
				}
				break
			} else if searchID <= 0 {
				fmt.Println("\nInvalid Entry")
				break
			} else {
				fmt.Println("\nNo Such Record found with that ID")
				break
			}
		}
	} else {
		fmt.Println("\nNo Records, Database is empty")
	}
}
func main() {
	fmt.Println("\nWELCOME TO ABC SCHOOL DIRECTORY ")
	menu()
}
