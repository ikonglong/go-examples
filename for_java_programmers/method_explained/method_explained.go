package main

import (
	"fmt"
	"time"
)

// https://yourbasic.org/golang/methods-explained/

type House struct {
	garage bool
}

func (p *House) HasGarage() bool { return p.garage }

type MyInt int

func (m MyInt) Positive() bool { return m > 0 }

func main() {
	house := new(House)
	fmt.Println(house.HasGarage()) // Prints "false" (zero value)

	// If you convert a value to a different type, the new value will have
	// the methods of the new type, but not the old.

	var m MyInt = 2
	m = m * m // The operators of the underlying type still apply.

	fmt.Println(m.Positive())        // Prints "true"
	fmt.Println(MyInt(3).Positive()) // Prints "true"

	var n int
	n = int(m) // The conversion is required.
	//n = m      // LEGAL
	fmt.Println("n:", n)

	// It’s idiomatic in Go to convert the type of an expression to access a specific method.

	var d int64 = 12345
	fmt.Println(d)                // 12345
	fmt.Println(time.Duration(d)) // 12.345µs
}
