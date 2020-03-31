package search

// Binary find given number on sort array
func Binary(sortArr []int, t int) bool {
	length := len(sortArr)
	for i := 0; i < length; i++ {
		mid := len(sortArr) / 2

		if sortArr[mid] == t {
			return true
		}

		if mid == 0 {
			return false
		}

		if sortArr[mid] > t {
			sortArr = sortArr[0:mid]
		} else {
			sortArr = sortArr[mid:]
		}
	}
	return false
}
