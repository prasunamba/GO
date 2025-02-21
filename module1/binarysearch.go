package module1

import "fmt"

func BinarySearch(list []int, num int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := low + (high-low)/2 // avoids int overflow
		fmt.Println(low, high, mid)
		if list[mid] == num {
			return mid
		} else if list[mid] > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
