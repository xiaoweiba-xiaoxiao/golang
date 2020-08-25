package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{IP: net.IPv4(192, 168, 234, 200), Port: 8080})
	if err != nil {
		fmt.Println("connection server failed, error:", err)
		return
	}
	defer socket.Close()
	myreader := bufio.NewReader(os.Stdin)
	for {
		data, _ := myreader.ReadString('\n')
		_, err := socket.Write([]byte(data))
		if err != nil {
			fmt.Println("send faild,Error,Error", err)
			continue
		}
		buf := make([]byte, 1024)
		n, rehost, err := socket.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("send faild,Error,Error", err)
			continue
		}
		fmt.Println("from", rehost, string(buf[:n]))
	}
}
