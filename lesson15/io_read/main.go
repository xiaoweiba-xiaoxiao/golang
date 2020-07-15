package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fs, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer fs.Close()
	var comtent []byte
	var buf [128]byte
	for {
		n, err := fs.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		//两个切片的append时，需要将第二个切片展开。
		comtent = append(comtent, buf[:n]...)
	}
	fmt.Printf(string(comtent))
}
