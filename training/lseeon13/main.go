package main

import (
	"fmt"
	"os"

	"xiao.com/golang/lesson/training/lesson12/student"
)

var stus *[]student.Student = student.Creatsu

func showMenu() {
	fmt.Printf("1. add student\n")
	fmt.Printf("2. modify student\n")
	fmt.Printf("3. show all student\n")
	fmt.Printf("4. exit\n")
}

func main() {

	for {
		stu := &student.Student{}
		showMenu()
		var sel int
		fmt.Scanf("%d", &sel)
		switch sel {
		case 1:
			stu.Add(stus)
		case 2:
			stu.Modify(stus)
		case 3:
			for _, val := range *stus {
				val.showMenu()
			}
		case 4:
			os.Exit(0)
		}
	}
}
