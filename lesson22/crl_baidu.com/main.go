package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("connection error,error:", err)
		return
	}
	defer conn.Close()
	data := "GET / HTTP/1.1\r\n"
	data += "Host: www.baidu.com\r\n"
	data += "Connection: close\r\n"
	data += "Cookie: BAIDUID=E0444FFA19A9704F9EF8D7E44E5EFC72:FG=1; BIDUPSID=E0444FFA19A9704F9EF8D7E44E5EFC72; PSTM=1558313678; COOKIE_SESSION=0_0_0_1_0_0_0_0_0_0_0_0_0_0_0_0_0_0_1597134380%7C1%230_0_1597134380%7C1; BD_UPN=12314753; BD_HOME=1; H_PS_PSSID=32662_1439_32620_31254_32351_32045_32117_32581"
	data += "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"
	data += "\r\n\r\n"
	_, err = io.WriteString(conn, data)
	if err != nil {
		fmt.Println("send data failed,error:", err)
		return
	}
	var buf []byte = make([]byte, 1024)
	for {
		n, err := conn.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("recv error,error:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
