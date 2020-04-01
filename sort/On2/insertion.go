package sort

// Insertion by right unsort array pick value, to insert left sort array
func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				unSortNum := arr[i]
				copy(arr[j+1:], arr[j:])
				arr[j] = unSortNum
			}
		}
	}
}
