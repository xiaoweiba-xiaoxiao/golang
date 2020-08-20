package main

import (
	"fmt"
	"time"
)

func worker(chnl chan int){
	for i := 0; i < 5; i++{		
		chnl <- i
		fmt.Printf("write %d\n",i)
	}
	close(chnl)
}

func main(){
	ch := make(chan int,2)
	go worker(ch)
	time.Sleep(2 * time.Second)
	for i := range ch{
		fmt.Printf("read value %d\n",i)
		time.Sleep(2 * time.Second)
	}	
}