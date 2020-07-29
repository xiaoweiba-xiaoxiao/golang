package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
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
	//获取传入的数据类型
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Struct {
		err = errors.New("plese pass Struct")
		return
	}
	//获取传入的值
	valinfo := reflect.ValueOf(data)
	//用于存储需要写入文件的数据
	var wrtdata []string
	//获取节点字节的信息
	for i := 0; i < t.NumField(); i++ {
		sectionField := t.Field(i)      //获取data的section信息
		sectinVal := valinfo.Field(i)   //获取section的值
		childField := sectionField.Type //获取section的类型 如果不是结构体就跳过
		if childField.Kind() != reflect.Struct {
			continue
		}
		//获取section的tag值
		tagVal := sectionField.Tag.Get("ini")
		if len(tagVal) == 0 {
			tagVal = sectionField.Name
		}
		//格式话section的节点名称
		section := fmt.Sprintf("[%s]\n", tagVal)
		wrtdata = append(wrtdata, section)
		//遍历每个节点的field信息
		for j := 0; j < childField.NumField(); j++ {
			childFieldType := childField.Field(j) //获取每个field的type
			childFieldVal := sectinVal.Field(j)   //获取每个field的val信息
			key := childFieldType.Tag.Get("ini")  // 获取每个field的tag信息

			if len(key) == 0 {
				key = childFieldType.Name
			}
			//格式化field信息添加到wrtdata里面
			fieldStr := fmt.Sprintf("%s=%v\n", key, childFieldVal)
			wrtdata = append(wrtdata, fieldStr)
		}
		wrtdata = append(wrtdata, "\n")
	}
	for _, val := range wrtdata {
		byteVal := []byte(val)
		result = append(result, byteVal...)
	}
	return
}

func WriteFile(file string, data []byte) (err error) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	newWrite := bufio.NewWriter(f)
	newWrite.Write(data)
	newWrite.Flush()
	return
}

func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("plese pass pointer")
		return
	}
	//获取指针位置的类型
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("plese pass struct")
		return
	}
	//存取节点字段
	var lastFieldName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}
		//解析节点
		if line[0] == '[' {
			lastFieldName, err = parseSection(line, index, typeStruct)
			if err != nil {
				return
			}
			fmt.Println(lastFieldName)
			continue
		}
		//解析节点所在的iteam
		err = parseItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%s line:%d", err, index+1)
			return
		}
	}
	return
}

//解析节点
func parseSection(line string, index int, typeStruct reflect.Type) (fieldName string, err error) {
	//节点非法输出,空括号
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error,invalid section:%s  line:%d", line, index+1)
		return
	}
	//节点非法情况2 "[" 没有 "]"
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error,invalid section:%s  line:%d", line, index+1)
		return
	}
	//左右都是括号的情况
	if line[0] == '[' && line[len(line)-1] == ']' {
		section := strings.TrimSpace(line[1 : len(line)-1])
		//左右括号中间为空的情况
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

//解析节点
func parseItem(lastFieldName string, line string, result interface{}) (err error) {
	//获取等号的位置
	equlindex := strings.Index(line, "=")
	//如果没有等号
	if equlindex == -1 {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}
	//获取等号左边的值为key,右边的值为val，并且去掉左右空格
	key := strings.TrimSpace(line[:equlindex])
	val := strings.TrimSpace(line[equlindex+1:])
	//如果key的长度为0，则key不存在
	if len(key) == 0 {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}
	//获取result的类型
	resultVal := reflect.ValueOf(result)
	//获取地址的filedname的值,值为一个结构体类型
	sectionVal := resultVal.Elem().FieldByName(lastFieldName)
	//获取值得类型
	sectionType := sectionVal.Type()
	//如果值得类型不是结构体 返回错误
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field must be struct:%s", line)
		return
	}
	//设置字段名
	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		//如果tag名称和key相同，将字段名赋值为当前的字段名
		if field.Tag.Get("ini") == key {
			keyFieldName = field.Name
			break
		}
	}
	//如果字段名字为空不处理则返回
	if len(keyFieldName) == 0 {
		return
	}
	//获取
	fieldval := sectionVal.FieldByName(keyFieldName)
	//如果获取不到值就返回错误
	if fieldval == reflect.ValueOf(nil) {
		return
	}
	// 根据字段值的类型来设置字段的值
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
