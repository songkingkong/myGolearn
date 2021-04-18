//数组按照指定次数反转
//解题逻辑
// reverse 先将整个数组反转
//反转次数和数组的原数个数求余得出实际转换次数
//根据余数将数组分成两部分，再执行反转即得出结果
package main

import "fmt"

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	rotate(list, 6)
	fmt.Println(list)
}
