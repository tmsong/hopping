/*
*
72. 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

思路：
用d[i][j]来表示A的前i个字母和B的前j个字母之间的编辑距离。
则有递推公式：
当A的第i个字符=B的第j个字符时，有d[i][j] = min(d[i][j-1]+1, d[i-1][j]+1, d[i-1][j-1])
当A的第i个字符!=B的第j个字符时，有d[i][j] = min(d[i][j-1]+1, d[i-1][j]+1, d[i-1][j-1]+1)
由此往下递推。
*/
package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 {
		return l2
	} else if l2 == 0 {
		return l1
	}
	//初始化distance二维数组，注意要多给一行用于表示word什么都不取的情况
	distance := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		distance[i] = make([]int, l2+1)
	}
	//再处理特殊的第一行和第一列
	for j := 1; j < l2+1; j++ {
		distance[0][j] = distance[0][j-1] + 1
	}
	for i := 1; i < l1+1; i++ {
		distance[i][0] = distance[i-1][0] + 1
	}
	//然后开始递推
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if word1[i-1] == word2[j-1] { //这里要注意，由于数组下标特性，word[i-1]才表示的是word的第i个字符
				distance[i][j] = min(distance[i][j-1]+1, distance[i-1][j]+1, distance[i-1][j-1])
			} else {
				distance[i][j] = 1 + min(distance[i][j-1], distance[i-1][j], distance[i-1][j-1])
			}
		}
	}
	return distance[l1][l2]
}

func min(i, j, k int) int {
	if i < j {
		if i < k {
			return i
		} else {
			return k
		}
	} else {
		if j < k {
			return j
		} else {
			return k
		}
	}
}

func main() {
	fmt.Println(minDistance("intention", "execution"))
}
