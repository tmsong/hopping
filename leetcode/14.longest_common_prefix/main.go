/*
*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

思路：
直接找肯定是最简单的。但还可以通过二分法来简化。
先确定最小的字符串长度，然后再比较此字符串长度上的字符，如果不匹配，则往下继续找。
*/
package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	minLength := math.MaxInt
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}
	if minLength == 0 {
		return ""
	}
	pos := (minLength - 1) / 2 //二分法
	//先往下找
	for {
		i := pos //初始化指针
	LOOP:
		for ; i >= 0; i-- {
			for j := 1; j < len(strs); j++ {
				if strs[j][i] != strs[j-1][i] { //逐一对比每位，如果出现不同，则重置指针位置
					break LOOP
				}
			}
		}
		if i >= 0 { //i没到底，表示对比失败了
			pos = i - 1  //重置pos，重新开始往下找
			if pos < 0 { //表示一点共同前缀都没有，直接不用找了
				return ""
			}
		} else { //对比成功，起码pos往前的前缀是共同的，那么从pos位置往后找
			i := pos + 1
			for ; i < minLength; i++ {
				for j := 1; j < len(strs); j++ {
					if strs[j][i] != strs[j-1][i] { //逐一对比每位，如果出现不同，则返回上一个符合的pos长度的前缀
						return strs[0][:i]
					}
				}
			}
			return strs[0][:i] //说明最大公共前缀就是最短的单词本身。
		}
	}
}

func main() {
	strs := []string{"flower", "flower", "flower", "flower"}
	fmt.Println(longestCommonPrefix(strs))
}
