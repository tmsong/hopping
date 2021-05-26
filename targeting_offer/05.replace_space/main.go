/*
剑指 Offer 05 替换空格

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."

思路：
strings.Builder

*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(s string) string {
	sb := strings.Builder{}
	for idx := range s {
		if s[idx] == ' ' {
			sb.Write([]byte{'%', '2', '0'})
		} else {
			sb.WriteByte(s[idx])
		}
	}
	return sb.String()
}
