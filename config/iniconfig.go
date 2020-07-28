package config

import (
	"errors"
	"fmt"
	"reflect"
	_ "reflect"
	"strconv"
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
	UserName string `ini:"user"`
	Passwd   string `ini:"passwd"`
	DataBase string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
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
		err = parseItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%s line:%d", err, index+1)
			return
		}
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

func parseItem(lastFieldName string, line string, result interface{}) (err error) {
	equlindex := strings.Index(line, "=")
	if equlindex == -1 {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}
	key := strings.TrimSpace(line[:equlindex])
	val := strings.TrimSpace(line[equlindex+1:])
	if len(key) == 0 {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}
	resultVal := reflect.ValueOf(result)
	sectionVal := resultVal.Elem().FieldByName(lastFieldName)
	sectionType := sectionVal.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field must be struct:%s", line)
		return
	}
	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		if field.Tag.Get("ini") == key {
			keyFieldName = field.Name
			break
		}
	}
	if len(keyFieldName) == 0 {
		return
	}
	fieldval := sectionVal.FieldByName(keyFieldName)
	if fieldval == reflect.ValueOf(nil) {
		return
	}
	switch fieldval.Type().Kind() {
	case reflect.String:
		fieldval.SetString(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, ok := strconv.ParseInt(val, 10, 64)
		if ok != nil {
			err = ok
			return
		}
		fieldval.SetInt(intVal)
	}
	return
}
