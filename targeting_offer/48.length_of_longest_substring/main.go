/*
*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

思路：用一个指针从左到右扫，然后用一个map来保存每个字符最后出现的位置。
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int, 0)
	max, left := 0, 0
	for right, c := range s {
		if lastLeft, ok := m[c]; ok && left <= lastLeft {
			left = lastLeft + 1
		}
		if right-left+1 > max {
			max = right - left + 1
		}
		m[c] = right //记录每个字符最后出现的位置
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
