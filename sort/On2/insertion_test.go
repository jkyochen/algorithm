package sort

import "testing"

func TestInsertion(t *testing.T) {
	arr := []int{4, 5, 7, 3, 9, 2, 3, 0, 5, 2, 10}
	Insertion(arr)
	if !isSort(arr) {
		t.Fatal("Insertion did not work as expected.")
	}
}
