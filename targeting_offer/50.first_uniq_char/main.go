/*
*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例 1:

输入：s = "abaccdeff"
输出：'b'
示例 2:

输入：s = ""
输出：' '

思路：只包含小写字母的话，维护一个26长度的数组，记录每个字符第一个出现的位置，如果重复出现，即抹去。
*/
package main

import "fmt"

func firstUniqChar(s string) byte {
	chars := make([]int, 26)
	for pos, b := range s {
		if chars[b-'a'] == 0 {
			chars[b-'a'] = pos + 1
		} else {
			chars[b-'a'] = -1
		}
	}
	min := -1
	for _, pos := range chars {
		if pos > 0 {
			if min == -1 {
				min = pos
			} else if pos < min {
				min = pos
			}
		}
	}
	if min == -1 {
		return ' '
	} else {
		return s[min-1]
	}
}

func main() {
	fmt.Println(string(firstUniqChar("")))
}
