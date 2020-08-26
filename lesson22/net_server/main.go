package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("listen port error, error:\n", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error,error", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {

	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read net stream error,error", err)
			break
		}
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
