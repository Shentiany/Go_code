package pointer_test

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var foo int
	var bar *int
	bar = &foo
	// %v 输出值	%T 输出类型
	*bar = 2
	// & 代表取地址
	fmt.Printf("%v %T \n", foo, foo)
	fmt.Printf("%v %T \n", *bar, bar)
}
