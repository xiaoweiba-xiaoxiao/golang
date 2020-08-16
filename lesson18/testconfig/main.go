package main

import (
	"fmt"

	"xiao.com/golang/lesson/config"
)

func main() {

	var conf config.ConfigInfo = config.ConfigInfo{
		Serverconf: config.ServerConfig{
			Ip:   "10.0.0.200",
			Port: 9200,
		},
		DBconfig: config.DBConfig{Host: "10.0.0.200", Port: 3306, UserName: "root", Passwd: "root", DataBase: "test"},
	}
	filename := "/home/golang/conf/test.ini"
	err := config.WriteFile(filename, conf)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	var myconfig config.ConfigInfo = config.ConfigInfo{}
	err = config.ReadConf(filename, &myconfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", myconfig)
}
