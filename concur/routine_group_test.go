package concur

import (
	"fmt"
	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

// lib doc: https://pkg.go.dev/golang.org/x/sync/errgroup

func TestLimitConcurrency(t *testing.T) {
	var workerPool errgroup.Group
	workerPool.SetLimit(4)
	counter := &atomic.Int32{}

	for i := 0; i <= 18; {
		workerPool.Go(func() error {
			counter.Add(1)
			fmt.Printf("counter: %v\n", counter.Load())
			time.Sleep(2 * time.Second)
			return nil
		})
		i++
	}
}
