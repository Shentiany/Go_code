package function_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	res := add(2, 3)
	fmt.Println(res)
	result, err := division(5, 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
func add(a int, b int) int {
	return a + b
}
func division(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ZeroError")
	}
	return a / b, nil
}
