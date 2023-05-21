package convert

import "testing"

type A struct {
	Name string
	Age  int
}

type B struct {
	Name string
}

func TestStructCopy(t *testing.T) {
	a := A{
		Name: "zjl",
		Age:  12,
	}
	var b B

	err := StructCopyFields(&b, a)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)
}
