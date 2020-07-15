package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "./test.tar.gz"
	//var r *bufio.Reader
	fz, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fz.Close()
	f, err := gzip.NewReader(fz)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	var content []byte
	var buf [128]byte
	for {
		n, err := f.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(content)
}
