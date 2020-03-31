package sort

// Bubble is Swap two adjacent numbers
func Bubble(arr []int) {
	swap := false
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap = true
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		if !swap {
			return
		}
	}
}
