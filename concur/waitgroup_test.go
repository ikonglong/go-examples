package concur

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// A sync.WaitGroup waits for a group of goroutines to finish.
// Note: a WaitGroup must not be copied after first use.

// Reference: https://yourbasic.org/golang/wait-for-goroutines-waitgroup/
func TestWaitGroupOfGoroutinesToFinish(t *testing.T) {
	var wg sync.WaitGroup
	// Call Add() to set the number of goroutines to wait for.
	wg.Add(2)
	go func() {
		// Do work.
		time.Sleep(6 * time.Second)
		fmt.Println("goroutine 1 finished")
		wg.Done()
	}()
	go func() {
		// Do work.
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine 2 finished")
		wg.Done()
	}()

	go func() {
		// Do work.
		time.Sleep(6 * time.Second)
		fmt.Println("goroutine 3 finished")
		wg.Done()
	}()

	// At the same time, Wait() is used to block until these two goroutines have finished.
	// 虽然 WaitGroup 对象只等两个协程执行完成，但协程 3 仍然可以执行。因为我们并未通过调用 WaitGroup 的
	// 某个方法来控制/阻塞协程的执行。事实上，WaitGroup 也没有这样的方法。
	wg.Wait()
	fmt.Println("The two goroutines have finished")
}

func TestUseWaitGroupToControlConcurrency(t *testing.T) {
	// WaitGroup 无法控制并发度，它不像 Java 中的信号量
}
