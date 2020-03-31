package search

import "testing"

func TestBinary(t *testing.T) {
	has := Binary([]int{1, 3, 5, 6, 7, 9, 33, 34}, 1)
	if !has {
		t.Fatal("Binary did not work as expected.")
	}

	has = Binary([]int{1, 3, 5, 6, 7, 9, 33, 34}, 34)
	if !has {
		t.Fatal("Binary did not work as expected.")
	}

	has = Binary([]int{1, 3, 5, 6, 7, 9, 33, 34}, 7)
	if !has {
		t.Fatal("Binary did not work as expected.")
	}

	has = Binary([]int{1, 3, 5, 6, 7, 9, 33, 34}, 6)
	if !has {
		t.Fatal("Binary did not work as expected.")
	}

	has = Binary([]int{1, 3, 5, 6, 7, 9, 33, 34}, 8)
	if has {
		t.Fatal("Binary did not work as expected.")
	}
}
