package main

import (
	_ "bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("xiaoxiao.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer file.Close()
	for i := 0; i < 10; i++ {
		file.WriteString("hello xiaoxiao!\n")
	}

}
