package main

import (
	"bufio"
	"fmt"
	"os"
)

func testBufio() {
	//var str_master string
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("plase input something:")
	str, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("you input: %s\n", str)
	}

}

func main() {
	testBufio()
}
