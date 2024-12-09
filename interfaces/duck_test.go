package interfaces

import (
	"fmt"
	"testing"
)

type iDuck interface {
	quack()
}

type kidMimickingDuck struct{}

func (k *kidMimickingDuck) quack() {
	fmt.Println("kid: 嘎嘎")
}

type duckToy struct{}

func (t *duckToy) quack() {
	fmt.Println("duck toy: 嘎嘎，嘎嘎")
}

func TestDuck(t *testing.T) {
	testDuck := func(d iDuck) {
		d = &duckToy{}
		d.quack()
	}
	testDuck(&kidMimickingDuck{})
}

func TestX(t *testing.T) {
	// type cast

	k := &kidMimickingDuck{}
	d := iDuck(k)
	d.quack()

	x(k)
}

func x(quacker any) {
	// type assertion
	type duck interface {
		quack()
	}
	switch quacker.(type) {
	case duck:
		// compile err: Cannot convert an expression of the type 'any' to the type 'duck'
		// q := duck(quacker)
		q := quacker.(duck)
		q.quack()
	}
}
