package main

import (
	"fmt"
)

func main() {
	m3 := make(map[string]int)
	if v, k := m3["pass"]; k {
		fmt.Println("key pass value is %d", v)
	} else {
		fmt.Println("key pass is not existing")
	}
	m3["pass"] = 0
	if v, k := m3["pass"]; k {
		fmt.Println("key pass value is %d", v)
	} else {
		fmt.Println("key pass is not existing")
	}
}
