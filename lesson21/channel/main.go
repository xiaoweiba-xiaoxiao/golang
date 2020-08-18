package main

func main() {
	var c chan int = make(chan int)
	c <- 100
	<-c
}
