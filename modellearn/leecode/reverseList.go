package main

import "fmt"

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
func rotate(nums []int, k int) {
	c := 4 % 7
	k %= len(nums)
	fmt.Println(c, k, len(nums))
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(list, 3)
	fmt.Println(list)
}
