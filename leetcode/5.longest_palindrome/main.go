/*
*
5. 最长回文子串

给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

思路：
1.使用一个动态规划的二维数组来递推，f[i][j]表示第i个到第j个是回文串， 如：babacabba 对应二维数组：

	b a b a c a b b a

b 1 0 1 0 0 0 0 0 0
a   1 0 1 0 0 0 0 0
b     1 0 0 0 1 0 0
a       1 0 1 0 0 0
c         1 0 0 0 0
a           1 0 0 1
b             1 1 0
b               1 0
a                 1

从右下角开始递推就行了。
*/
package main

import "fmt"

func longestPalindrome(s string) string {
	bs := []byte(s)
	f := make([][]bool, len(bs))
	for i := 0; i < len(bs); i++ {
		f[i] = make([]bool, len(bs))
	}
	var max, maxI, maxJ int //分别记录最长的回文长度以及此时的i和j值
	for i := len(bs) - 1; i >= 0; i-- {
		for j := i; j < len(bs); j++ {
			if j == i { //j=i的时候，相当于初始化斜对角上的矩阵值
				f[i][j] = true
			} else if j-i == 1 { //j和i相邻，这时如果相等也是对称的
				if bs[i] == bs[j] {
					f[i][j] = true
				}
			} else if f[i+1][j-1] == true && bs[i] == bs[j] { //这时j至少比i大2的情况
				f[i][j] = true
			}
			if f[i][j] == true && j-i+1 > max { //更新max
				max, maxI, maxJ = j-i+1, i, j
			}
		}
	}
	return string(bs[maxI : maxJ+1])
}

func main() {
	fmt.Println(longestPalindrome("babacabba"))
}
