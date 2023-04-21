// 最小公倍数
package main

import (
	"fmt"
)

func gcd(a, b int) int {
	tem := a % b
	if tem > 0 {
		return gcd(b, tem)
	} else {
		return b
	}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	var a int
	var b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Printf("%d", lcm(a, b))
}
