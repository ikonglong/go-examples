package structcopy

import "testing"

// goverter:converter
// goverter:output:file ./converter_impl.go
// goverter:output:package github.com/ikonglong/go-examples/structcopy
//
//go:generate go run github.com/jmattheis/goverter/cmd/goverter@v1.3.1 gen ./
type Converter interface {
	ConvertItems(source []Input) []Output

	// goverter:ignore Irrelevant
	// goverter:map Nested.AgeInYears Age
	Convert(source Input) Output
}

type Input struct {
	Name   string
	Nested InputNested
}

type InputNested struct {
	AgeInYears int
}

type Output struct {
	Name       string
	Age        int
	Irrelevant bool
}

func TestUseGoverter(t *testing.T) {

}
