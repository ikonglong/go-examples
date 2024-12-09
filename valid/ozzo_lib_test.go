package valid

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"regexp"
	"testing"
)

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// State cannot be empty, and must be a string consisting of five digits
		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}

func TestOzzo(t *testing.T) {
	// data := "example"
	// err := validation.Validate(data,
	// 	validation.Required,       // not empty
	// 	validation.Length(5, 100), // length between 5 and 100
	// 	is.URL,                    // is a valid URL
	// )
	// fmt.Println(err)
	// // Output:
	// // must be a valid URL

	// a := Address{
	// 	Street: "123",
	// 	City:   "Unknown",
	// 	State:  "Virginia",
	// 	Zip:    "12345",
	// }
	//
	// err := a.Validate()
	// fmt.Println(err)
	// // Output:
	// // Street: the length must be between 5 and 50; State: must be in a valid format.

	c := map[string]interface{}{
		"Name":  "Qiang Xue",
		"Email": "q",
		"Address": map[string]interface{}{
			"Street": "123",
			"City":   "Unknown",
			"State":  "Virginia",
			"Zip":    "12345",
		},
	}

	err := validation.Validate(c,
		validation.Map(
			// Name cannot be empty, and the length must be between 5 and 20.
			validation.Key("Name", validation.Required, validation.Length(5, 20)),
			// Email cannot be empty and should be in a valid email format.
			validation.Key("Email", validation.Required, is.Email),
			// Validate Address using its own validation rules
			validation.Key("Address", validation.Map(
				// Street cannot be empty, and the length must between 5 and 50
				validation.Key("Street", validation.Required, validation.Length(5, 50)),
				// City cannot be empty, and the length must between 5 and 50
				validation.Key("City", validation.Required, validation.Length(5, 50)),
				// State cannot be empty, and must be a string consisting of two letters in upper case
				validation.Key("State", validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
				// State cannot be empty, and must be a string consisting of five digits
				validation.Key("Zip", validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
			)),
		),
	)
	fmt.Println(err)
	// Output:
	// Address: (State: must be in a valid format; Street: the length must be between 5 and 50.); Email: must be a valid email address.
}
