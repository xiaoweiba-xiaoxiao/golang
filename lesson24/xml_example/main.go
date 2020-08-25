package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Servers struct {
	Name    xml.Name `xml:"Servers"`
	Version string   `xml:"version,attr"`
	Server  []Server `xml:"server"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIp   string `xml:"serverIp"`
}

func main() {
	data, err := ioutil.ReadFile("./config.xml")
	if err != nil {
		fmt.Println("open config.xml faild,error:", err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println("unmarshal faild,error", err)
		return
	}
	fmt.Printf("%#v", servers)

}
