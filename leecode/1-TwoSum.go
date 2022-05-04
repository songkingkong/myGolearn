package twosum
#一个数组，一个值，取数组内两个数和等于值的两个数的下标

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
func twosum() {
	list := []int{4, 5, 6, 7, 8}
	a := twoSum(list, 15)
	fmt.Println(a)
}
