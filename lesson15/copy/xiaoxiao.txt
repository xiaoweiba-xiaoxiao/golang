package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer srcFile.Close()
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer dstFile.Close()
	io.Copy(dstFile, srcFile)
}

func main() {
	src, dst := "./main.go", "./xiaoxiao.txt"
	copyFile(src, dst)
}
