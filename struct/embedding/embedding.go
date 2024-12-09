package embedding

import "github.com/ikonglong/go-examples/struct/embedded"

type addr embedded.Address

type User struct {
	addr
	Name string
}

// -company-prefixes github.com/ikonglong -project-name github.com/ikonglong/$ProjectName$ -rm-unused $FilePath$
