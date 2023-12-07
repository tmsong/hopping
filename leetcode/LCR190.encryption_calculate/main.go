/*
*
计算机安全专家正在开发一款高度安全的加密通信软件，需要在进行数据传输时对数据进行加密和解密操作。假定 dataA 和 dataB 分别为随机抽样的两次通信的数据量：

正数为发送量
负数为接受量
0 为数据遗失

请不使用四则运算符的情况下实现一个函数计算两次通信的数据量之和（三种情况均需被统计），以确保在数据传输过程中的高安全性和保密性。

思路：用位运算实现加法。

n=dataA⊕dataB	非进位和：异或运算
c=dataA & dataB<<1	进位：与运算+左移一位


（和 s ）=（非进位和 n ）+（进位 c ）。即可将
s = dataA + dataB 转化为：

s=dataA+dataB⇒s=n+c
循环求 n 和 c ，直至进位 c=0  ；此时 s=n，返回 n 即可。
*/
package main

func encryptionCalculate(dataA int, dataB int) int {
	for dataB != 0 {
		c := (dataA & dataB) << 1
		dataA = dataA ^ dataB
		dataB = c
	}
	return dataA
}

func main() {

}
