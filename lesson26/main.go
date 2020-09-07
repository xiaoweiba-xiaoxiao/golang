package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func testString() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	_, err = conn.Do("set", "abc", 100)
	if err != nil {
		fmt.Printf("set failed,error:%v\n", err)
		return
	}
	r, err := redis.Int(conn.Do("get", "abc"))
	if err != nil {
		fmt.Printf("get failed,error:%v\n", err)
		return
	}
	fmt.Printf("the key get %v\n", r)
}

func testHash() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	_, err = conn.Do("HSet", "book", "abc", 100)
	if err != nil {
		fmt.Printf("hset failed,error:%v\n", err)
		return
	}
	r, err := redis.Int(conn.Do("HGet", "book", "abc"))
	if err != nil {
		fmt.Printf("get failed,error:%v\n", err)
		return
	}
	fmt.Printf("the key get %v\n", r)

}

func testList() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	_, err = conn.Do("Lpush", "books", "hehe", "haha", "xixi")
	if err != nil {
		fmt.Printf("hset failed,error:%v\n", err)
		return
	}
	r, err := redis.Strings(conn.Do("Lrange", "books", 0, -1))
	if err != nil {
		fmt.Printf("get failed,error:%v\n", err)
		return
	}
	fmt.Printf("the key get %v\n", r)

}

func main() {
	// testHash()
	testList()
}
