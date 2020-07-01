package main

import (
	"fmt"
	"time"
)

func example() (int64, string) {
	now := time.Now()
	timestamp := now.Unix()
	return timestamp, fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func testTicker() {
	ticker := time.Tick(10 * time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		func() {
			fmt.Println("done")
		}()
	}
}

func testConst() {
	fmt.Printf("nano second:%d\n", time.Nanosecond)
	fmt.Printf("micrp second:%d\n", time.Microsecond)
	fmt.Printf("mili second:%d\n", time.Millisecond)
	fmt.Printf("second:%d\n", time.Second)
}

func testFormat() {
	now := time.Now()
	timeStr := now.Format("2006/01/02 15:04:05")
	fmt.Printf("time:%s\n", timeStr)
}

func main() {
	// timestamp, timestr := example()
	// fmt.Println(timestamp)
	// fmt.Println(timestr)
	// testTicker()
	// testConst()
	testFormat()
}
