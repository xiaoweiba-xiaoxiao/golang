package main

import (
	_ "encoding/json"
	"fmt"
	"os"

	"xiao.com/golang/lesson/training/lesson12/student"
)

var students *[]student.Student = student.Createstu()

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

func cheakStudent(name string) (bool, int) {
	for index, val := range *students {
		if val.Name == name {
			return true, index
		}
	}
	return false, -1
}

func addStudent() {
	stu := inputStudent()
	if err, index := cheakStudent(stu.Name); !err {
		*students = append(*students, *stu)
	} else {
		(*students)[index] = *stu
	}
	fmt.Println("user add success")
}

func modifyStudent() {
	var modify_name string
	fmt.Println("please enter student's name that you want to modify:")
	fmt.Scanf("%s\n", &modify_name)
	if err, index := cheakStudent(modify_name); err {
		stu := inputStudent()
		(*students)[index] = *stu
	} else {
		fmt.Printf("student %s is not exised", modify_name)
	}
}

func showStudent() {
	fmt.Printf("%#v\n", students)
}

func main() {

	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			addStudent()
		case 2:
			modifyStudent()
		case 3:
			showStudent()
		case 4:
			os.Exit(0)
		}
	}
}
