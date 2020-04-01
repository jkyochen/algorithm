package sort

import (
	"algorithm/sort/helper"
	"testing"
)

func TestQuick(t *testing.T) {
	arr := []int{4, 5, 7, 3, 9, 2, 3, 0, 5, 2, 10}
	Quick(arr)
	if !helper.IsSort(arr) {
		t.Fatal("Quick did not work as expected.")
	}
}
