package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"xiao.com/golang/lesson/lesson24/yaml"
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
	file := "test.yaml"
	yaml1 := yaml.NewYaml()
	data, _ = yaml1.ReadFile(file)
	fmt.Println(data)
	ok := yaml1.WriteFile(file)
	fmt.Println(ok)
}
