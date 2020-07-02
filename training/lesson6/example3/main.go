package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "~!@#$%^&"
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l 用户密码长度")
	flag.StringVar(&charset, "t", "num",
		`-t 用户密码生成字符集.
	num:纯数字.
	char:[a-z][A-Z].
	mix:字母和数字.
	advance:使用数字,字母,特殊符号`)
	flag.Parse()
}

func genneratePasswd() string {
	var passwd []byte = make([]byte, length, length)
	var sourceStr string
	switch charset {
	case "char":
		sourceStr = CharStr
	case "mix":
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	case "advance":
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	default:
		sourceStr = NumStr
	}
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()
	//fmt.Printf("length:%d \n", length)
	//fmt.Printf("charset:%s\n", charset)
	passwd := genneratePasswd()

	fmt.Println(passwd)
}
