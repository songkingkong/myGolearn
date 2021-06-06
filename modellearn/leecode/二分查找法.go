package main

import "fmt"

var list = [...]int{1,2,3,4,5,6,7}
var v = 2

func erFen(key int, listArg [7]int) {
	for i := range listArg{
		fmt.Println(i,key)
	}
}
func main()  {
	erFen(v,list)
}