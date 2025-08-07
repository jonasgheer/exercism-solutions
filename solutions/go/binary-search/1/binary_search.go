package binarysearch

func SearchInts(list []int, key int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		middle := (low + high) / 2
		if key < list[middle] {
			high = middle - 1
		} else if key > list[middle] {
			low = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
