package json

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"

	"testing"
)

type book struct {
	Title     string `json:"title"`
	Authors   []string
	Publisher publisher
}

type publisher struct {
	Name    string
	Address string
}

func TestUnmarshalToMap(t *testing.T) {
	var jsonStr = `
	{
      "title": "三体",
      "authors": ["刘慈欣"],
	  "publisher": {
		"name": "火星人",
		"address": "火星村 10 号"
      }
    }
	`

	var objMap map[string]interface{}

	// 报错：json: Unmarshal(non-pointer map[string]interface {})
	// err := json.Unmarshal([]byte(jsonStr), objMap)

	// the following code is OK
	err := json.Unmarshal([]byte(jsonStr), &objMap)
	fmt.Printf("objMap: %v\n", objMap)

	// objMap = map[string]interface{}{}
	// err := json.Unmarshal([]byte(jsonStr), &objMap)
	// fmt.Printf("objMap: %v\n", objMap)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshalNonMapJsonToMap(t *testing.T) {
	var arrayJson = `["a", "b", "c"]`
	var objMap map[string]interface{}
	err := json.Unmarshal([]byte(arrayJson), &objMap)
	assert.Nil(t, err)
}
