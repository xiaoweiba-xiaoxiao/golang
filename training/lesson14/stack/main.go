package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"github.com/urfave/cli"
)

func getInput() (string,error){
	inputRead := bufio.NewReader(os.Stdin)
	input,err := inputRead.ReadString('\n')
	return input,err
}

func transPost(express string) (postexpress []string,err error){
	var opstack stack
	var i int
LABLE:
	for i < len(express) {
		switch {
		case express[i]>='0' && express[i] <='9':
			var number []byte
			for ;i<len(express);i++{
				if express[i]<'0' || express[i] >'9'{
					break
				}
				number = append(number,express[i])			
			}
			postexpress = append(postexpress,string(number))
		case express[i] == '+' || express[i] == '-' || express[i] == '*' || express[i] == '/':
			if opstack.empty() {
				opstack.push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABLE
			}
			data,_ := opstack.Top()
			if data[0] == '(' || data[0]==')'{
				 opstack.push(fmt.Sprintf("%c",express[i]))
				 i++
				 continue LABLE
			}
			if (express[i] == '+' || express[i] == '-') || ((express[i] == '*' || express[i] == '/') && (data[0] == '*' || data[0] == '/')){
				postexpress = append(postexpress,data)
				opstack.pop()
				i++
				continue LABLE
			}
		case express[i] == '(':
			opstack.push(fmt.Sprintf("%c",express[i]))
			i++
			continue LABLE
		case express[i] == ')':
			for !opstack.empty(){
				data,_ := opstack.pop()
				if data[0]=='(' {
					break
				}
				postexpress = append(postexpress,data)
			}
			i++
		default:
			err = fmt.Errorf("invalid express:%v", express[i])
			return
		}
	}
	for !opstack.empty() {
		data, _ := opstack.pop()
		if data[0] == '#' {
			break
		}
		postexpress = append(postexpress, data)
		//numStack.Push(data)
	}
	return
}

func cacl(postexpress string) (result int,err error){
	
	return
}

func process(c *cli.Context) (err error){
	for {
		express,_ := getInput()
		if len(express) == 0 {
			continue
		}
		express = strings.TrimSpace(express)
		postexpress,errRest := transPost(express)
		if errRest != nil{
			err = errRest
			fmt.Println(err)
			return
		}
		result,errRes := cacl(postexpress)
		if errRes !=nil {
			err = errRes
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

func main(){
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error{
		return process(c)
	}
	app.Run(os.Args)
}