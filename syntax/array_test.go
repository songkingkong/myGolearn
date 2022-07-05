/*
数组类型的应用
*/
package syntax

import (
	"fmt"
	"testing"
)

/*
声明数组arr1里面有两个参数，分别是1，2
*/
var arr1 = [2]int{1, 2}

/*
声明数组arr2，“...”代表参数的个数根据赋值的数量来
*/
var arr2 = [...]int{1, 2, 3, 4, 5}

/*
声明数组arr3，有两个参数，没有赋值，且赋值必须在函数内进行(如果中括号内没有数字则声明的是切片)
*/
var arr3 = [2]int{}

/*
声明多层数组arr4包含两个数组，每个数组有两个参数
*/
var arr4 = [2][2]int{{1, 2}, {3, 4}}

/*
获取数组内参数的方式：arr1[0]
*/

func TestArray(t *testing.T) {
	arr3[0] = 6
	arr3[1] = 7
	fmt.Println(arr1, arr3)
	/*通过自增下标的方式轮询数组 */
	for i := 0; i < len(arr2); i++ {
		fmt.Println(i, arr2[i])
	}
	/*通过range 遍历数组 */
	for _, a := range arr4 {
		fmt.Println(a)
	}
}
