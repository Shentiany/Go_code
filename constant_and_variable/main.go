package main

import "fmt"

// 常量的声明
const pi = 3.14
const e = 2.7

const (
	a = 1
	b = 2
)

const (
	//n2, n3 默认为 n1 的值
	n1 = 10
	n2
	n3
)

const (
	x1 = iota //0
	x2 = iota //1
	x3 = iota //2
)

const (
	y1 = iota
	y2 = iota
	_
	y4 = iota
)

func main() {
	// 标准变量声明
	var age int = 18
	var name string = "small tiger"
	fmt.Println(age, name)

	// 支持变量推导
	var age0 = 18
	var name0 = "small tiger"
	fmt.Println(age0, name0)

	// 批量变量声明
	var (
		age1  = 19
		name1 = "big tiger"
	)
	fmt.Println(age1, name1)

	//短变量声明, 不能再函数外使用
	name2 := "tiger"
	age2 := 20
	fmt.Println(age2, name2)

	fmt.Println(n1, n2, n3)
	fmt.Println(x1, x2, x3)
	fmt.Println(y1, y2, y4)
}
