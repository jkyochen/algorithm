package sort

import (
	"algorithm/sort/helper"
	"testing"
)

func TestSelection(t *testing.T) {
	arr := []int{4, 5, 7, 3, 9, 2, 3, 0, 5, 2, 10}
	Selection(arr)
	if !helper.IsSort(arr) {
		t.Fatal("Selection did not work as expected.")
	}
}
