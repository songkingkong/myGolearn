package syntax

import (
	"fmt"
	"testing"
	"time"
)

func slowFunc(b chan string) {
	time.Sleep(time.Second * 2)
	b <- "sleep finished"
}
func TestGoroutine(t *testing.T) {
	c := make(chan string)
	go slowFunc(c)
	msg := <-c
	fmt.Printf(msg, "\n")
}
