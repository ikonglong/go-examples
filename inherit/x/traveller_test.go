package x

import (
	"fmt"
	"testing"
)

type ITraveller interface {
	carryOutTravel()
	makePlan() string
	prepareEssentials() string
	travel()
}

type TravellerBase struct {
	ITraveller
}

func (t *TravellerBase) carryOutTravel() {
	fmt.Println(t.makePlan())
	fmt.Println(t.prepareEssentials())
	t.ITraveller.travel()
	// t.travel()
}

func (t *TravellerBase) makePlan() string {
	return "TravellerBase: make a plan"
}

func (t *TravellerBase) prepareEssentials() string {
	return "TravellerBase: prepare essentials"
}

type Traveller struct {
	ITraveller
}

func (t *Traveller) travel() {
	fmt.Println("an Traveller is travelling")
}

func TestTraveller(t *testing.T) {
	tr := &Traveller{
		ITraveller: &TravellerBase{},
	}
	base := &TravellerBase{
		ITraveller: tr,
	}
	tr.ITraveller = base
	tr.carryOutTravel()
}
