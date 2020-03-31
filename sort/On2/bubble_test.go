package sort

import "testing"

func TestBubble(t *testing.T) {
	arr := []int{4, 5, 7, 3, 9, 2, 3, 0, 5, 2, 10}
	Bubble(arr)
	if !isSort(arr) {
		t.Fatal("Bubble did not work as expected.")
	}
}
