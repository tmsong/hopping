/*
1049. 最后一块石头的重量 II

有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

示例 1：

输入：stones = [2,7,4,1,8,1]
输出：1
解释：
组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。

经过归纳法证明（此处参考leetcode网站）
最后可以转换为0-1背包问题，也就是在总和为sum的数组中，找出一定的数，
使其重量总和neg在不超过sum/2的情况下尽可能大。也就是越接近sum/2越好。

这里定义二维数组dp，其中dp[i+1][j]表示前i个石头是否能凑出重量j。所以dp[0][]为不选任何石头的状态。
对于第i个石头，考虑凑出重量j：
若j<stones[i]，则不能选第i个石头，则有dp[i+1][j] = dp[i][j]
若j>stones[i]，存在选或不选两个决策，
不选时有 dp[i+1][j]=dp[i][j]，选时需要考虑能否凑出重量j-stones[i]，即dp[i+1][j]=dp[i][j−stones[i]]。
若二者均为假则dp[i+1][j]为假，否则 dp[i+1][j] 为真。
求出 dp[n][]后，在所有为真的dp[n][j]中，最大的 j 即为 neg能取到的最大值。
代入 sum−2⋅neg中即得到最后一块石头的最小重量。

比如 2，4，1 sum = 7
dp =
1 0 0 0 0 0 0 0
1 0 1 0 0 0 0 0
1 0 1 0 1 0 1 0
1 1 1 1 1 1 1 1

（0-1背包问题）
*/
package main

import (
	"fmt"
	"math"
)

func lastStoneWeightII(stones []int) int {
	var sum int
	for _, stone := range stones {
		sum += stone
	}
	prev, cur := make([]bool, sum+1), make([]bool, sum+1) //虽然是动态规划，但是背包问题只需要两行
	prev[0] = true
	for i := 0; i < len(stones); i++ {
		for j := 0; j < sum+1; j++ {
			if prev[j] { //对应dp[i][j] = true的情况
				cur[j] = true
			} else if j >= stones[i] && prev[j-stones[i]] { //对应dp[i][j-stone[i]] = true的情况
				cur[j] = true
			}
		}
		prev, cur = cur, prev //将cur赋值给prev，这样不用一个矩阵，只需要两行即可搞定
	}
	min := math.MaxInt
	for j := 0; j <= sum/2; j++ {
		if prev[j] && sum-2*j < min {
			min = sum - 2*j
		}
		if min == 0 {
			return min
		}
	}
	return min
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}
