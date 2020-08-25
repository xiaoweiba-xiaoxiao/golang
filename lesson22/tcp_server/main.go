package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error,err:", err)
			break
		}
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error,err:", err)
			continue
		}
		go process(conn)
	}
}
