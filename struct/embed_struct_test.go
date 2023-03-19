package struct_ex

import (
	"fmt"
	"testing"
)

type Animal struct {
	name string
}

func (a *Animal) Eat() {
	fmt.Printf("%v is eating\n", a.name)
}

type Dog struct {
	Animal
}

func TestCreateSubType(t *testing.T) {
	dog := Dog{
		Animal{
			name: "Snoopy",
		},
	}
	dog.Eat()
	fmt.Printf("dog name: %v\n", dog.name)
}
