package condition__test

import "testing"

func TestIfMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("It is not 0~3")
		}
	}
}

func TestIfMultiCase1(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 != 0:
			t.Log("Odd")
		default:
			t.Log("Unknown")
		}
	}
}
