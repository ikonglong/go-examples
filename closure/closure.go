package main

import (
	"fmt"
	"sync"
)

// New returns a function Count.
// Count prints the number of times it has been invoked.
func New() (Count func()) {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

type Customer struct {
	name string
	pay  int
}

func raisePay(c Customer, wg *sync.WaitGroup) {
	fmt.Printf("[raisePay] addr of c: %p\n", &c)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		c.pay = 500
		go func() {
			fmt.Printf("[closure i=%v] addr of c: %p, pay of c: %v\n", i, &c, c.pay)
			wg.Done()
		}()
	}
}

func main() {
	// 每一次调用 New，就创建了新的变量 n。因此，f1 和 f2 互不影响
	f1, f2 := New(), New()
	f1() // 1
	f2() // 1 (different n)
	f1() // 2
	f2() // 2

	wg := sync.WaitGroup{}
	c := Customer{
		name: "Alice",
		pay:  100,
	}
	fmt.Printf("before raisePay, addr of c: %p, c.pay: %v\n", &c, c.pay)
	raisePay(c, &wg)
	wg.Wait()
}
