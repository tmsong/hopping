/**
剑指 Offer 20. 表示数值的字符串

请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。

数值（按顺序）可以分成以下几个部分：

若干空格
一个小数或者整数
（可选）一个'e'或'E'，后面跟着一个整数
若干空格
小数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
下述格式之一：
至少一位数字，后面跟着一个点 '.'
至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
一个点 '.' ，后面跟着至少一位数字
整数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
至少一位数字
部分数值列举如下：

["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
部分非数值列举如下：

["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]

思路：按照题目的描述，稍微用一点递归，需要注意非常多的特殊情况

*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"+100", "5e2", "-123", "3.1416", "-1E-16", "0123"}
	for _, s := range strs {
		fmt.Println(isNumber(s))
	}
}

func isNumber(s string) bool {
	str := strings.TrimSpace(strings.ToLower(s))
	if strings.Contains(str, "e") {
		strs := strings.SplitN(str, "e", 2)
		return isDecimals(strs[0]) && isInteger(strs[1])
	} else {
		return isDecimals(str)
	}
}

func isDecimals(str string) bool {
	if len(str) == 0 {
		return false
	}
	if str[0] == '+' || str[0] == '-' {
		str = str[1:]
	}
	if len(str) == 0 || str == "." {
		return false
	}
	strs := strings.SplitN(str, ".", 2)
	for _, s := range strs {
		if len(s) > 0 {
			for _, b := range s {
				if b < '0' || b > '9' {
					return false
				}
			}
		}
	}
	return true
}

func isInteger(str string) bool {
	if len(str) == 0 {
		return false
	}
	if str[0] == '+' || str[0] == '-' {
		str = str[1:]
	}
	if len(str) == 0 {
		return false
	}
	for _, b := range str {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}
