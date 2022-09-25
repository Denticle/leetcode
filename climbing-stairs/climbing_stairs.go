package main

import "fmt"

func main() {
	climbStairs(2)
}

func climbStairs(n int) int {
	// 初始情况，p记录n-2的值，q记录n-1的值，r记录n的值
	p, q, r := 0, 0, 1
	// 如果 n 有值，那么n为 n-2 的值加上 n -1的值
	// 逆向推导，从第三个数开始，每个数都为前两个数的和
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	// 此时结果即为最终结果
	fmt.Println(r)
	return r
	//dp[i] 代表有多少种爬楼梯的方式
	//递推公式
	// dp[i] = dp[i-1] + dp[i-2]
	// if n <= 1 {
	// 	return n
	// }
	// // 初始化
	// dp := []int{1, 2}
	// // 遍历 ： 从前向后遍历
	// // 举例子
	// // 1 2 3 5 8 13 21
	// // 优化点：循环从第三个数开始
	// for i := 2; i < n; i++ {
	// 	sum := dp[0] + dp[1]
	// 	dp[0] = dp[1]
	// 	// 优化点：只保留两个数，在最终时会再统计到数组中
	// 	dp[1] = sum
	// }
	// return dp[1]
}
