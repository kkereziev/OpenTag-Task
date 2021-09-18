package helpers

func BinarySearch(arr []rune, letter rune) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < letter {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != letter {
		return false
	}

	return true
}
