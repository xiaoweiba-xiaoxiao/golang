package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	file := "./test.tar.gz"
	var r *bufio.Reader
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
	r = bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done read file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
	// var content []byte
	// var buf [128]byte
	// for {
	// 	n, err := f.Read(buf[:])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("Error: %s\n", err)
	// 		return
	// 	}
	// 	content = append(content, buf[:n]...)
	// }
	// fmt.Println(string(content))
}
