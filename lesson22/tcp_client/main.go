package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println("connection Error,Error: ", err)
		return
	}
	defer conn.Close()
	readers := bufio.NewReader(os.Stdin)
	for {
		data, err := readers.ReadString('\n')
		if err != nil {
			fmt.Println("read input error, err:", err)
			break
		}
		stream := []byte(data)
		_, err = conn.Write(stream)
		if err != nil {
			fmt.Println("write error, err:", err)
			break
		}
	}
}
