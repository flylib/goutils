package structfun

import (
	"fmt"
	"testing"
)

type A struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
type B struct {
	Age  int64
	Name string
}

func TestStructCopy(t *testing.T) {
	a := A{
		Name: "小明",
		Age:  16,
	}

	var b B
	StructCopy(&b, a)
	fmt.Println(b) //{16 小明}
}
