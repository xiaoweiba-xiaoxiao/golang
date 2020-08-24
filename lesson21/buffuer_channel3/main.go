package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	Id, Number int
}

type Result struct {
	job    Job
	resNum int
}

var jobChan chan Job = make(chan Job, 10)
var resChan chan Result = make(chan Result, 10)

func digist(num int) int {
	var sum int
	for num != 0 {
		dig := num % 10
		sum += dig
		num /= 10
	}
	return sum
}

/*创建job*/
func CreatJob() {
	for i := 0; i < cap(jobChan); i++ {
		num := rand.Intn(999)
		job := Job{Id: i, Number: num}
		jobChan <- job
	}
	close(jobChan)
}

/*添加工作*/
func worker(wg *sync.WaitGroup) {
	for job := range jobChan {
		result := Result{job, digist(job.Number)}
		resChan <- result
	}
	wg.Done()

}

/*打印输出*/
func PrintResult(done chan bool) {
	/*for res := range resChan {
		fmt.Printf("result of %d job.Number is %d is %d\n", res.job.Id, res.job.Number, res.resNum)
	}*/
	res := <-resChan
	fmt.Printf("result of %d job.Number is %d is %d\n", res.job.Id, res.job.Number, res.resNum)
	done <- true
}

/*创建工作池 */
func createWorkerPool() {
	var wg sync.WaitGroup
	for i := 0; i < cap(resChan); i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(resChan)
}

func main() {
	rand.Seed(time.Now().Unix())
	done := make(chan bool)
	go CreatJob()
	go PrintResult(done)
	go createWorkerPool()
	<-done
}
