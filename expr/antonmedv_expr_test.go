package main

import (
	"fmt"
	"github.com/antonmedv/expr"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestSimpleAddExpr(t *testing.T) {
	program, err := expr.Compile(`2 + 2`)
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, nil)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 4, output.(int))
}

func TestMapAsEnv_PassSomeVarsToExpr(t *testing.T) {
	env := map[string]any{
		"foo": 100,
		"bar": 200,
	}

	program, err := expr.Compile(`foo + bar`, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 300, output.(int))
}

func TestMapAsEnv_InferTypeOfExpr_CheckItAgainstEnv(t *testing.T) {
	env := map[string]any{
		"name": "Anton",
		"age":  35,
	}

	_, err := expr.Compile(`name + age`, expr.Env(env))
	if err != nil {
		panic(err) // Will panic with "invalid operation: string + int"
	}
}

func TestMapAsEnv_UseFuncContainedInMap(t *testing.T) {
	env := map[string]interface{}{
		"greet":   "Hello, %v!",
		"names":   []string{"world", "you"},
		"sprintf": fmt.Sprintf,
	}

	code := `sprintf(greet, names[0])`

	program, err := expr.Compile(code, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println("output:", output)
	assert.Equal(t, "Hello, world!", output.(string))
}

type PostList struct {
	Posts []Post `expr:"posts"`
}

func (pl PostList) Format(t time.Time) string {
	return t.Format(time.RFC822)
}

func (pl PostList) TellNumOfPosts() string {
	return fmt.Sprintf("%d posts", len(pl.Posts))
}

type Post struct {
	Body string
	Date time.Time
}

func TestStructAsEnv_UseFuncDefinedOnStruct_RenameFieldWithExprTag(t *testing.T) {
	// code := `map(posts, Format(#.Date) + ": " + #.Body + "; ")`
	code := `map(posts, Format(.Date) + ": " + .Body + "; ")`
	env := PostList{
		Posts: []Post{
			{"Oh My God!", time.Now()},
			{"How you doing?", time.Now()},
			{"Could I be wearing any more clothes?", time.Now()},
		},
	}
	program, err := expr.Compile(code, expr.Env(PostList{}))
	if err != nil {
		panic(err)
	}
	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

	code = `TellNumOfPosts()`
	program, err = expr.Compile(code, expr.Env(PostList{}))
	if err != nil {
		panic(err)
	}
	output, err = expr.Run(program, env)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func TestUseCustomFunc(t *testing.T) {
	type book struct {
		title string
	}

	toJSON := func(params ...any) (any, error) {
		bytes, err := json.Marshal(params[0])
		return string(bytes), err
	}

	env := map[string]interface{}{
		"book": book{title: "Test in Action"},
	}

	code := `toJSONStr(book)`
	code = `book | toJSONStr()`
	program, err := expr.Compile(code, expr.Env(map[string]interface{}{"book": ""}), expr.Function("toJSONStr", toJSON))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}

type OSSClient struct{}

func (cli OSSClient) GetFile(fileKey string) string {
	return "input.json"
}

type GetFileExt struct {
	theFunc func(params ...any) (any, error)
}

var ossClient = OSSClient{}

func TestXxx(t *testing.T) {
	code := `fileExt("file_key_1")`
	funcHolder := GetFileExt{
		theFunc: func(params ...any) (any, error) {
			fileKey := params[0].(string)
			// fileName := ossClient.GetFile(fileKey)
			// strings.Cut()strings.IndexAny(".")
			return ossClient.GetFile(fileKey), nil
		},
	}
	program, err := expr.Compile(code, expr.Function("fileExt", funcHolder.theFunc))
	if err != nil {
		panic(err)
	}
	output, err := expr.Run(program, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
