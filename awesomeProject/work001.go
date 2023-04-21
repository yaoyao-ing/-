// 作业  819. 递归求阶乘
package main

import "fmt"

func plus(num int) int {
	if num < 2 {
		return num
	}
	return plus(num-1) * num
}
func main() {
	var num int
	fmt.Scanf("%d", &num)
	fmt.Printf("%d", plus(num))
}
