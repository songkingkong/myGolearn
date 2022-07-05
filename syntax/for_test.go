package syntax

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	/* for循环内 k 等于下标，v等于值 */
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	/* for循环内 使用_表示丢弃下标 */
	for _, v := range m1 {
		fmt.Println(v)
	}
}
