package main

import "fmt"

/*
len()可以用来查看数组或slice的长度
cap()可以用来查看数组或slice的容量,在容量不足时增长是乘以2
*/
/*
切片的声明方式 make(值 类型，参数的数量)
*/
var s0 = make([]string, 2)

/*
切片的声明时没有标明参数的个数，如果带有个数那声明的是list
*/
var s1 []int
var s2 = []string{"eee", "fff"}

func main() {
	s0[0] = "aaa"
	s0[1] = "bbb"
	s0 = append(s0, "ccc")
	s1 = append(s1, 1, 2)
	s2 = append(s2, "ggg")
	fmt.Println(s0, s1, s2)
	arr := []int{2, 3, 5, 7, 11, 13}
	sli := arr[1:4]
	fmt.Println(sli)
	fmt.Println(len(sli))
	fmt.Println(cap(sli))
}
