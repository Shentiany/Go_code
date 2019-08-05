package struct_test

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	p := person{
		name: "jack",
		age:  18,
	}
	fmt.Println(p.age)
	fmt.Println(p.name)
}
