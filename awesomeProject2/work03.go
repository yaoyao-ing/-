// 722. 数字序列和它的和
package main

import "fmt"

func main() {
	var m, n, sum int
	for {
		sum = 0
		fmt.Scanf("%d %d", &m, &n)
		if m == 0 || n == 0 {
			break
		}
		if m < n {
			m, n = n, m
		}
		for i := n; i <= m; i++ {
			fmt.Printf("%d ", i)
			sum += i
		}
		fmt.Println("Sum=", sum)
	}
}
