package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	for _, v := range m1 {
		fmt.Println(v)
	}
}
