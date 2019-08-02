package string_test

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	/**
	byte 代表UTF-8字符串的当个字节的值
	rune代表单个Unicode字符
	*/
	var s string = `hello world\n`  // hello world\n
	var s1 string = "hello world\n" // hello world
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Printf("%c\n", s[0])

	// 字符串的拼接，只能拼接字符串
	foo := "hello" + "world"
	fmt.Println(foo)

	//字符串的遍历
	x := "hello world 你好 世界"
	for i, n := 0, len(x); i < n; i++ {
		var ch2 byte = x[i]
		fmt.Printf("%d-%c ", i, ch2)
	}
}
