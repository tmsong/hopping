/**
剑指 Offer 38. 字符串的排列

输入一个字符串，打印出该字符串中字符的所有排列。

你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

思路：
1.回溯：也就是DFS，为了防止字符串内有重复字符，可以先将字符串排序，然后每次取字符串，都保证从相同字符的最左边一个取起。
2.下一个排列：由于字符有大小关系，可以先生成字符串的最小排列，然后快速得到字典序中下一个更大的排列。

*/
package main

import "sort"

func main() {

}

func permutation(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	n := len(t)
	perm := make([]byte, 0, n)
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			ans = append(ans, string(perm))
			return
		}
		for j, b := range vis {
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, t[j])
			backtrack(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return
}

///////思路2，下一个排列（更优的思路）

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func nextPermutation(nums []byte) bool {
	n := len(nums)
	i := n - 2 //从最后开始往前找，找到第一个值比右边小的元素i
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for j >= 0 && nums[i] >= nums[j] { //从右往左找到第一个比此元素i大的元素j
		j--
	}
	nums[i], nums[j] = nums[j], nums[i] //交换i，j位置。
	reverse(nums[i+1:])                 //逆向所有原本i位置右边的数（将逆向变成了顺向）
	return true
}

func permutation2(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	for {
		ans = append(ans, string(t))
		if !nextPermutation(t) {
			break
		}
	}
	return
}
