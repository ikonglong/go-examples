package struct_ex

import (
	"fmt"
	"testing"
)

type Quacker interface {
	quack()
}

type ToyDuck struct {
	name string
}

func (d *ToyDuck) quack() {
	fmt.Printf("In ToyDuck.quack()...\n")
	fmt.Printf("%v quacks\n", d.name)
}

// DuckPen 鸭子笔
type DuckPen struct{}

func (p *DuckPen) draw() {
	fmt.Printf("duck-pen draws a line\n")
}

func TestCallToNilToyDuck(t *testing.T) {
	var toy *ToyDuck
	toy.quack()
}

func TestCallToNilDuckPen(t *testing.T) {
	var p *DuckPen
	p.draw()
}
