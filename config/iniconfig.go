package config

import (
	"errors"
	"fmt"
	"reflect"
	_ "reflect"
	"strings"
)

type ConfigInfo struct {
	Serverconf ServerConfig `ini:"server"`
	DBconfig   DBConfig     `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type DBConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini: "port"`
	UserName string `ini: "username"`
	Passwd   string `ini: "password"`
	DataBase string `ini: database`
}

func Marshal(data interface{}) (result []byte, err error) {
	// 	v := reflect.ValueOf(data)
	// 	t := v.Type()
	// 	k := t.Kind()
	// 	switch k {
	// 		case
	// 	}
	return
}

func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	_ = lineArr
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("plese pass pointer")
		return
	}
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("plese pass struct")
		return
	}
	var lastFieldName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}
		if line[0] == '[' {
			lastFieldName, err = parseSection(line, index, typeStruct)
			if err != nil {
				return
			}
			fmt.Println(lastFieldName)
			continue
		}
		equlindex := strings.Index(line,"=")
		if equlindex == -1 {
			err = fmt.Errorf("syntax error,invalid section:%s  line:%d", line, index+1)
			return
		}
		tagname := line[:equlindex]
		value := line[equlindex+1:]
		for i := 0;i < 
		
	}
	return
}

func parseSection(line string, index int, typeStruct reflect.Type) (fieldName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error,invalid section:%s  line:%d", line, index+1)
		return
	}
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error,invalid section:%s  line:%d", line, index+1)
		return
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		section := strings.TrimSpace(line[1 : len(line)-1])
		if len(section) == 0 {
			err = fmt.Errorf("syntax error,invalid section:%s  line:%d", line, index+1)
			return
		}
		for i := 0; i < typeStruct.NumField(); i++ {
			field := typeStruct.Field(i)
			tag := field.Tag.Get("ini")
			if tag == section {
				fieldName = field.Name
				break
			}
		}
	}
	return
}
