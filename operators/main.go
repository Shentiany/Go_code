package main

import "fmt"

func main() {
	// 算术运算符
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)

	// 关系运算符
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	// 逻辑运算符
	fmt.Println(10 > 4 && 5 < 4)
	fmt.Println(!(10 > 4) || 5 < 4)

	// 位运算符
	x := 1 // 用二进制表示为001
	y := 5 // 用二进制表示为101
	// 两位均为1才为1
	fmt.Println(x & y) //1
	// 两位中有一个为1就为1
	fmt.Println(x | y) //5
	// 两位不一样则为1
	fmt.Println(x ^ y) //4
	// a<<b 把a的各二进位全部左移b位
	fmt.Println(y << 2)
	// a>>b 把a的各二进位全部右移b位
	fmt.Println(y >> 2)

	// 赋值运算符
	var n int
	a = 5
	a += 5
	fmt.Println(n)
}
