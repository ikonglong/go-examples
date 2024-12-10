package select_test

import (
	"fmt"
	"testing"
	"time"
)

// ref: https://yourbasic.org/golang/select-explained/

// The select statement waits for multiple send or receive operations simultaneously.
//
// - The statement blocks as a whole until one of the operations becomes unblocked.
// - If several cases can proceed, a single one of them will be chosen at random.
// - The select statement is not a loop statement, it only executes once.

func TestSelectCommonUsageMistake1(t *testing.T) {
	// 报错：fatal error: all goroutines are asleep - deadlock!
	// why? 从理论上说，select 语句可能阻塞。但从逻辑上说，select 语句不会阻塞。
	ch1 := make(chan int)
	ch1 <- 1
	select {
	case r := <-ch1:
		fmt.Printf("received from ch1: %d\n", r)
	}
}

func TestSelectCommonMistake2(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	defer close(ch1)
	defer close(ch2)

	// 注意，可能会报告错误：panic: send on closed channel
	go func() {
		ch1 <- 1
		fmt.Println("send 1 to ch1")
	}()
	go func() {
		ch2 <- 2
		fmt.Println("send 2 to ch2")
	}()

	time.Sleep(3 * time.Second)

	select {
	case r := <-ch1:
		fmt.Printf("received from ch1: %d\n", r)
	case r := <-ch2:
		fmt.Printf("received from ch2: %d\n", r)
	}
}

func TestSelectBasicUsage(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer close(ch1)
	defer close(ch2)

	go func() {
		for i := 0; ; i++ {
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := 0; ; i++ {
			ch2 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case r := <-ch1:
			fmt.Printf("received from ch1: %d\n", r)
		case r := <-ch2:
			fmt.Printf("received from ch2: %d\n", r)
		}
	}
}

func TestSelectWithNilChanCase(t *testing.T) {
	var ch1 chan int
	ch2 := make(chan int)
	defer close(ch2)

	go func() {
		for i := 0; ; i++ {
			ch2 <- i
		}
	}()

	time.Sleep(3 * time.Second)

	// Send and receive operations on a nil channel block forever.
	// This can be used to disable a channel in a select statement:
	select {
	case <-ch1:
		fmt.Println("received from ch1")
	case r := <-ch2:
		fmt.Printf("received from ch2: %d\n", r)
	}
}

func TestSelectWithDefaultCase(t *testing.T) {
	var ch chan string

	// The default case is always able to proceed
	// and runs if all other cases are blocked.

	// never blocks
	select {
	case x := <-ch: // Send and receive operations on a nil channel block forever.
		fmt.Println("Received", x)
	default:
		fmt.Println("Nothing available")
	}
}

func TestSelectWithZeroCases(t *testing.T) {
	go func() {
		for {
			fmt.Println(" a heartbeat")
			time.Sleep(5 * time.Second)
		}
	}()

	// A select statement blocks until at least one of its cases can proceed.
	// With zero cases this will never happen.
	//
	// A typical use would be at the end of the main function in some
	// multithreaded programs. When main returns, the program exits and does
	// not wait for other goroutines to complete.
	select {}
}

func TestSelectBlockingOpWithTimeout(t *testing.T) {
	AFP := make(chan string)
	defer func() {
		close(AFP)
		fmt.Println("closed AFP")
	}()

	// The function time.After is part of the standard library; it waits for
	// a specified time to elapse and then sends the current time on the
	// returned channel.
	select {
	case news := <-AFP:
		fmt.Println(news)
	case <-time.After(30 * time.Second):
		fmt.Println("Time out: no news in one minute")
	}
}

// As a toy example you can use the random choice among cases
// that can proceed to generate random bits.
func TestGenInfiniteRandBinarySeq(t *testing.T) {
	rand := make(chan int)
	defer close(rand)

	// 从输出的二进制序列是随机的可知，select 确实是随机选择 case 的。
	go func() {
		for {
			select {
			case bit := <-rand:
				fmt.Println(bit)
			}
		}
	}()

	for {
		select {
		case rand <- 0:
		case rand <- 1:
		}
	}
}
