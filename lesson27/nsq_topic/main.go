package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strings"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func initProducer(str string) error {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	nsqAddress := "localhost:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	read := bufio.NewReader(os.Stdin)
	for {
		data, err := read.ReadString('\n')
		if err != nil {
			fmt.Printf("read string failed,err:%v\n", err)
			continue
		}
		if data == "stop" {
			break
		}
		err = producer.Publish("user_info", []byte(data))
		if err != nil {
			fmt.Printf("publish message failed,err:%v\n", err)
			continue
		}
		fmt.Printf("publish data:%s succ\n", data)
	}
}
