/*
*
424. 替换后的最长重复字符
给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
在执行上述操作后，返回 包含相同字母的最长子字符串的长度。

示例 1：
输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。

示例 2：
输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。
可能存在其他的方法来得到同样的结果。

思路：
用左右指针来表示。然后用一个哈希表来记录左右指针之间，每个值出现的次数。
不断移动右指针来增加长度，如果出现（左右之间的长度-出现最多字符的次数）> k，就移动左指针，并更新哈希表中的次数统计。
*/
package main

func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
