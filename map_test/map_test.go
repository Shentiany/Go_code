package map_test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	foo := map[int]string{1: "hello", 2: "world"}
	fmt.Println(foo)
	emptyMap := map[int]int{}
	fmt.Println(emptyMap)
	var emptyMap2 map[string]string
	fmt.Println(emptyMap2)
	bar := make(map[string]string, 10)
	bar["name"] = "jack"
	bar["age"] = "8"
	fmt.Println(bar)
	delete(bar, "name")
	fmt.Println(bar)
}
