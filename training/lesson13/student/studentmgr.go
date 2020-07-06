package student

import "fmt"

type Studentmgr struct {
	Allstudent []*Student
}

func (s *Studentmgr) Add(t *Student) (err error) {
	for index, stu := range s.Allstudent {
		if stu.Name == t.Name {
			s.Allstudent[index] = t
			fmt.Printf("student %s add success\n", t.Name)
			return
		}
	}
	s.Allstudent = append(s.Allstudent, t)
	fmt.Printf("student %s add success\n", t.Name)
	return
}

func (s *Studentmgr) Modify(t *Student) (err error) {
	for index, stu := range s.Allstudent {
		if stu.Name == t.Name {
			s.Allstudent[index] = t
			fmt.Printf("student %s modify success\n", t.Name)
			return
		}
	}
	return fmt.Errorf("student %s not exist\n", t.Name)
}

func (s *Studentmgr) Show() {
	for _, stu := range s.Allstudent {
		fmt.Printf("%#v\n", stu)
	}
}
