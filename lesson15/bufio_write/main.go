package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./xiaoxiao.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer file.Close()
	newWrite := bufio.NewWriter(file)
	str := "hello xiaoxiao!\n"
	for i := 0; i < 10; i++ {
		newWrite.Write([]byte(str))
	}
	newWrite.Flush()
}
