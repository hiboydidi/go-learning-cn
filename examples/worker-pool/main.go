package main

import (
	"fmt"
	"sync"
)

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		results <- Result{
			JobID: job,
			Value: job * job,
		}
		fmt.Printf("Worker %d 完成任务 %d\n", id, job)
	}
}

func main() {
	const workerCount = 3
	const jobCount = 10

	jobs := make(chan int)
	results := make(chan Result, jobCount)

	var wg sync.WaitGroup
	for id := 1; id <= workerCount; id++ {
		wg.Add(1)
		go worker(id, jobs, results, &wg)
	}

	go func() {
		for jobID := 1; jobID <= jobCount; jobID++ {
			jobs <- jobID
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("任务 %d 的结果是 %d\n", result.JobID, result.Value)
	}
}
