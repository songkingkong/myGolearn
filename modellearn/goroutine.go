package main

import (
	"fmt"
	"time"
)

func slowFunc(b chan string) {
	time.Sleep(time.Second * 2)
	b <- "sleep finished"
}
func main() {
	c := make(chan string)
	go slowFunc(c)
	msg := <-c
	fmt.Printf(msg)
}
