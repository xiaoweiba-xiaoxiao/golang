package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 8080})
	if err != nil {
		fmt.Println("listen port failed, error:", err)
		return
	}
	defer server.Close()
	for {
		buf := make([]byte, 1024)
		n, remoteAppdrr, err := server.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("read data faild,error:", err)
			continue
		}
		fmt.Println("from", remoteAppdrr, "message:", string(buf[:n]))
		senddata := []byte("hello client")
		_, err = server.WriteToUDP(senddata, remoteAppdrr)
		if err != nil {
			fmt.Println("发送数据失败!", err)
			return
		}
	}
}
