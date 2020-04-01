package sort

// Quick Through the partition point, put the number greater than the partition point to the right of the partition point, and the number smaller than the partition point to the left
func Quick(arr []int) {
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, start, end int) {
	if start >= end {
		return
	}
	point := partition(arr, start, end)
	quick(arr, start, point-1)
	quick(arr, point+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
