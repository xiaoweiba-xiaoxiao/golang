package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Println("input:")
	str, err := reader.ReadString('\n')
	if err == nil {
		fmt.Fprintf(os.Stdout, "you input is %s", str)
	}
}
