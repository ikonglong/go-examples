package _defer

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type user struct {
	name string
}

func (u user) hello() {
	fmt.Printf("[user.hello] address of u: %p, value: %+v\n", &u, u)
	fmt.Printf("hello, I'm %s\n", u.name)
}

func (u *user) hello2() {
	fmt.Printf("[user.hello2] address of u: %p, value: %+v\n", u, u)
	fmt.Printf("hello, I'm %s\n", u.name)
}

func TestDefer(t *testing.T) {
	funcToDeferCall := func(x *user) {
		fmt.Printf("[func to defer call] address of x: %p, value: %+v\n", x, x)
	}
	u := user{
		name: "peter",
	}
	fmt.Printf("address of u: %p, value: %+v\n", &u, u)
	defer funcToDeferCall(&u)
}

func TestDefer2(t *testing.T) {
	u := user{
		name: "peter",
	}
	fmt.Printf("address of u: %p, value: %+v\n", &u, u)
	defer u.hello()
}

func TestDefer3(t *testing.T) {
	u := &user{
		name: "peter",
	}
	fmt.Printf("address of u: %p, value: %+v\n", u, u)
	defer u.hello()
}

func TestDefer4(t *testing.T) {
	u := &user{
		name: "peter",
	}
	fmt.Printf("address of u: %p, value: %+v\n", u, u)
	defer u.hello2()
}

func TestDefer5(t *testing.T) {
	u := user{
		name: "peter",
	}
	fmt.Printf("address of u: %p, value: %+v\n", &u, u)
	defer u.hello2()
}

func bigSlowOperation() {
	// defer trace("bigSlowOperation")() // don't forget the extra parentheses
	defer trace("bigSlowOperation")
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
