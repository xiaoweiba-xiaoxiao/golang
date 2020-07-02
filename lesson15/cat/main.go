package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		content, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", content)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("%s\n", flag.Arg(i))
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}
		r := bufio.NewReader(file)
		cat(r)
	}
}
