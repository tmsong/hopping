/**
剑指 Offer 19. 正则表达式匹配

请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释:因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释:".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释:因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

思路：分都空/一个空/两个都不空的情况讨论
其中，两个都不空分为：pattern的第二位是不是* 来讨论
*/
package main

import "fmt"

func main() {
	fmt.Println(isMatch("aab", "c*a*b"))
}

func isMatch(s string, p string) bool {
	return isMatchBytes([]byte(s), []byte(p))
}

func isMatchBytes(sb []byte, pb []byte) bool {
	//case 1 都是空，返回true
	if len(sb) == 0 && len(pb) == 0 {
		return true
	}
	//case 2 string不为空，pattern为空，返回false
	if len(sb) > 0 && len(pb) == 0 {
		return false
	}
	//case 3 string为空，pattern不为空，判断pattern是否是a*b*c*d*这种模式
	if len(sb) == 0 && len(pb) != 0 {
		if len(pb)%2 != 0 {
			return false
		}
		for i := 1; i < len(pb); i += 2 {
			if pb[i] != '*' {
				return false
			}
		}
		return true
	}
	//case 4 都不为空的情况
	//case 4.1 如果pattern不够两位，或pattern的第二位不是*，按常规处理
	if len(pb) == 1 || pb[1] != '*' {
		if pb[0] == sb[0] || pb[0] == '.' {
			return isMatchBytes(sb[1:], pb[1:])
		}
		return false
	}
	//case 4.2 如果pattern的第二位是*的情况
	var maxEqual int //记录string开头最长的与pattern[0]相等的字符长度
	for i := 0; i < len(sb); i++ {
		if sb[i] == pb[0] || pb[0] == '.' {
			maxEqual++
		} else {
			break
		}
	}
	//分别去掉这些与pattern[0]相等的字符，将其与pattern[2:]比较
	for i := 0; i <= maxEqual; i++ {
		if isMatchBytes(sb[i:], pb[2:]) {
			return true
		}
	}
	return false
}
