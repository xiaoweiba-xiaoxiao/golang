package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {
}

func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func main() {
	err := initConumer("user_info", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed,err:%v\n", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}

func initConumer(topic string, channel string, address string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return err
	}
	consumer := &Consumer{}
	c.AddHandler(consumer)

	if err := c.ConnectToNSQLookupd(address); err != nil {
		return err
	}
	return nil
}
