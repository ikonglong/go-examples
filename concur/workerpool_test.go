package concur

import (
	"fmt"
	"testing"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func TestWorkerPool(t *testing.T) {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动 3 个 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 提交 5 个 job
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		v := <-results
		fmt.Printf("result: %v\n", v)
	}
}
