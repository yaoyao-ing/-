// 765. 字符串加空格
package main

import (
	"bufio"
	"fmt"
	"os"

	//"os"
	"strings"
)

func main() {
	var str string

	reader := bufio.NewReader(os.Stdin) // 标准输入输出
	str, _ = reader.ReadString('\n')    // 回车结束
	str = strings.TrimSpace(str)        // 去除最后一个空格
	//fmt.Printf(str)
	var res []string
	for _, s := range str {
		res = append(res, string(s))
	}
	newStr := strings.Join(res, " ")
	fmt.Println(newStr)
}
