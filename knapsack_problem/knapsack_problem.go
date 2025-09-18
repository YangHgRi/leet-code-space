package knapsack_problem

import "fmt"

// max 辅助函数，返回两个整数中较大的一个
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// packing 函数计算给定物品价值、重量和背包最大承重下的最大总价值
func packing(values, weights []int, maxWeight int) int {
	// dp 数组用于存储在特定背包容量下的最大价值
	// dp[j] 表示容量为 j 时能获得的最大价值
	// 长度为 maxWeight + 1，dp[0] 到 dp[maxWeight]
	dp := make([]int, maxWeight+1)

	// 遍历每一个物品
	for i := 0; i < len(values); i++ {
		w := weights[i] // 当前物品的重量
		v := values[i]  // 当前物品的价值

		// 从最大容量 maxWeight 倒序遍历到当前物品的重量 w
		// 倒序遍历确保每个物品只被选择一次 (0/1 背包的关键)
		for j := maxWeight; j >= w; j-- {
			// 对于当前容量 j，考虑两种情况：
			// 1. 不放入当前物品：价值保持为 dp[j] (即不考虑当前物品时的最大价值)
			// 2. 放入当前物品：价值为 v (当前物品价值) + dp[j-w] (剩余容量 j-w 的最大价值)
			// 取这两种情况中的最大值
			dp[j] = max(dp[j], v+dp[j-w])
		}
	}

	// 循环结束后，dp[maxWeight] 即为背包最大承重下的最大总价值
	return dp[maxWeight]
}

func main() {
	m := 10 // 背包最大承重

	values := []int{5, 5, 2, 3, 4, 2}  // 物品价值
	weights := []int{7, 4, 3, 2, 3, 1} // 物品重量

	fmt.Println(packing(values, weights, m)) // 输出最大总价值
}
