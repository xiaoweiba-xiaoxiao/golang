package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func main() {
	pool := newPool("localhost:6379", "")
	for {
		time.Sleep(time.Second)
		conn := pool.Get()
		conn.Do("set", "abc", 100)
		r, err := redis.Int(conn.Do("get", "abc"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v\n", r)
	}
}
