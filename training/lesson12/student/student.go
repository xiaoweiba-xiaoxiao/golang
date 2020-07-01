package student

import "fmt"

type Student struct {
	Name  string  `json:"name"`
	Sex   string  `json:"name"`
	Age   int     `josn:"age"`
	Grade string  `json:"grade"`
	Score float32 `json:"score"`
}

func New(name, sex string, age int, grade string, score float32) (stu *Student) {
	stu = &Student{
		Name:  name,
		Sex:   sex,
		Age:   age,
		Grade: grade,
		Score: score,
	}
	return
}

func Createstu() *[]Student {
	stu := make([]Student, 0)
	return &stu
}

func input() *Student {
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
	stu := New(name, sex, age, grade, score)
	return stu
}

func check(name *string, stus *[]Student) (bool, int) {
	for index, val := range *stus {
		if val.Name == *name {
			return true, index
		}
	}
	return false, -1
}

func (s *Student) Add(stus *[]Student) {
	if s == nil {
		s = input()
	}
	if err, index := check(&s.Name, stus); !err {
		*stus = append(*stus, *s)
	} else {
		(*stus)[index] = *s
	}

}

func (s *Student) Modify(stus *[]Student) {
	if s == nil {
		fmt.Println("the student that will modify is nil")
		return
	}
	if err, index := check(&s.Name, stus); !err {
		fmt.Print("student %s is not exised", s.Name)
	} else {
		stu := input()
		(*stus)[index] = *stu
	}

}

func (stu *Student) Show() {
	fmt.Printf("%#v\n", stu)
}
