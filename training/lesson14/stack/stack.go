package main

import (
	"fmt"
)

type stack struct{
	data [1024]string
	top int
}

func (s *stack)push(d string) {
	s.data[s.top] = d
	s.top ++
}

func (s *stack)pop() (ret string,err error){
	if s.top == 0 {
		err = fmt.Errorf("stack is empty")
		return
	}
	s.top--
	ret = s.data[s.top]
	return
}

func (s *stack)Top() (ret string,err error){
	if s.top == 0 {
		err = fmt.Errorf("stack is empty")
		return
	}
	ret = s.data[s.top -1 ]
	return
}

func (s *stack)empty() bool{
	return s.top == 0
}