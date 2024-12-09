package inherit

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"

	"testing"
)

type ITraveller interface {
	carryOutTravel()
	makePlan() string
	prepareEssentials() string
	travel()
}

type TravellerGrandpa struct {
	ITraveller
}

func (t *TravellerGrandpa) carryOutTravel() {
	fmt.Println(t.makePlan())
	fmt.Println(t.prepareEssentials())
	t.travel()
}

func (t *TravellerGrandpa) makePlan() string {
	return "a travel plan made by a TravellerGrandpa object"
}

type TravellerFather struct {
	*TravellerGrandpa
	Name *string
}

func (t *TravellerFather) prepareEssentials() string {
	n := *t.Name
	return fmt.Sprintf("the essentials prepared by a TravellerFather object{name=%s}", n)
}

type Traveller struct {
	*TravellerFather
}

func (t *Traveller) travel() {
	fmt.Println("an Traveller is travelling")
}

func TestTravellerBaseLevel1(t *testing.T) {
	convey.Convey("Given type TravellerGrandpa partially implements interface ITraveller", t, func() {
		convey.Convey("When create an object of type TravellerGrandpa", func() {
			obj := &TravellerGrandpa{}
			convey.Convey("Then ", func() {
				convey.So(obj.prepareEssentials, convey.ShouldBeNil)
			})
		})
	})
}

func TestTraveller(t *testing.T) {
	tr := Traveller{
		TravellerFather: &TravellerFather{
			TravellerGrandpa: &TravellerGrandpa{
				ITraveller: &Traveller{},
			},
		},
	}
	tr.carryOutTravel()
}

func TestTraveller2(t *testing.T) {
	name := "Bob"
	tr := Traveller{
		TravellerFather: &TravellerFather{
			Name:             &name,
			TravellerGrandpa: &TravellerGrandpa{},
		},
	}
	tr.TravellerGrandpa.ITraveller = &tr
	tr.carryOutTravel()
}
