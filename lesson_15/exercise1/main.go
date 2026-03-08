package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main()  {
	fmt.Println("Lesson 15")
	ctx := context.Background()
	jobs := make(chan Job, 100)
	for i := range 100 {
		jobs <- Job{ID: strconv.Itoa(i), Payload: fmt.Sprintf("task-%d", i)}
	}
	close(jobs)

	res := WorkerPool(ctx, 5, jobs)
	for r := range res {
		fmt.Printf("JobID: %s | Output: %s | Err: %v\n", r.JobId, r.Output, r.Err)
	}
}

type Job struct {
	ID string
	Payload string
}

type Result struct {
	JobId string
	Output string
	Err error
}

func processResult(job Job) Result{
	time.Sleep(10 * time.Millisecond)
	return Result{
		JobId: job.ID,
		Output: fmt.Sprintf("processed: %s", job.Payload),
	}
}

func WorkerPool(ctx context.Context, n int, jobs <-chan Job) <-chan Result{
	results := make(chan Result, len(jobs))

	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				select{
				case <- ctx.Done():
					return 
				default:
					results <- processResult(job)
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return  results
}