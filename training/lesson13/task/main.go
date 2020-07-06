package main

import (
	"fmt"
	"os"

	"xiao.com/golang/lesson/training/lesson13/student"
)

var (
	stumgr = &student.Studentmgr{}
)

func showMenu() {
	fmt.Printf("1. add student\n")
	fmt.Printf("2. modify student\n")
	fmt.Printf("3. show all student\n")
	fmt.Printf("4. exit\n")
}

func inputStudent() *student.Student {
	var (
		name  string
		sex   string
		age   int
		grade string
		score float32
	)
	fmt.Println("please enter student's name:")
	fmt.Scanf("%s\n", &name)
	fmt.Println("please enter student's sex:")
	fmt.Scanf("%s\n", &sex)
	fmt.Println("please enter student's age:")
	fmt.Scanf("%d\n", &age)
	fmt.Println("please enter student's grade:")
	fmt.Scanf("%s\n", &grade)
	fmt.Println("please enter student's score:")
	fmt.Scanf("%f\n", &score)
	stu := student.New(name, sex, age, grade, score)
	return stu
}

func main() {

	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			stu := inputStudent()
			stumgr.Add(stu)
		case 2:
			stu := inputStudent()
			stumgr.Modify(stu)
		case 3:
			stumgr.Show()
		case 4:
			os.Exit(0)
		}
	}
}
