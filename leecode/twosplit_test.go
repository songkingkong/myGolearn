package leecode

import (
	"fmt"
	"testing"
)

var list = []int{1, 2, 3, 4, 5, 6, 7}
var v = 2

func binarySearch(arr []int, k int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (l + r) / 2
		if k == arr[mid] {
			return mid
		}
		if k < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
func TestTwosplit(t *testing.T) {
	a := binarySearch(list, v)
	fmt.Println(a)
}
