package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr4 := [4]string{1: "hello", 3: "world"}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3, arr4)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for index, x := range arr3 {
		t.Log(index, x)
	}
	for _, x := range arr3 {
		t.Log(x)
	}
}
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec := arr3[:2]
	t.Log(arr3Sec)
}
func TestArrayGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	// 以下操作会改变原有的值
	Q2[0] = "Unknown"
	t.Log(year)
}
func TestSliceArrayComparing(t *testing.T) {
	// array可以比较大小， 而slice不可以比较大小
	a := [4]int{1, 2, 3, 4}
	b := [4]int{1, 2, 3, 4}
	if a == b {
		t.Log("equal")
	}
}
