package config

import (
	_ "fmt"
	"testing"
)

// func TestConfig(t *testing.T) {
// 	configFile := "config.ini"
// 	data, err := ioutil.ReadFile(configFile)
// 	if err != nil {
// 		t.Error("read file failed")
// 		return
// 	}
// 	var result *ConfigInfo = &ConfigInfo{}
// 	//var testint *int
// 	err = UnMarshal(data, result)
// 	if err != nil {
// 		t.Error("unmashal filed Error:", err)
// 	}
// 	t.Logf("unmashal success,conf: %#v", result)
// }

func TestWriteConfig(t *testing.T) {
	data := ConfigInfo{
		Serverconf: ServerConfig{Ip: "192.168.234.200", Port: 100},
		DBconfig:   DBConfig{Host: "192.168.234.200", Port: 3306, UserName: "root", Passwd: "root", DataBase: "test"},
	}
	result, err := Marshal(data)
	if err != nil {
		t.Error("mashal failed Error:", err)
	}
	err = WriteFile("/home/golang/conf/test.conf", result)
	if err != nil {
		t.Errorf("write file faild %s", err)
	}
	t.Logf("write success")
}
