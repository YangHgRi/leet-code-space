/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=30204
 *
 * [121] 买卖股票的最佳时机
 */

// @lcpr-template-start
package leet_code_space

import "math"

// @lcpr-template-end
// @lc code=start

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0 // 如果价格列表为空，利润为0
	}

	maxProfit := 0
	// 初始化 minPrice 为一个足够大的值，确保第一个价格能更新它
	// 或者直接用第一个价格初始化，但需注意循环起始
	minPrice := math.MaxInt32

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // 遇到更低的价格，更新最低买入点
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice // 计算当前价格卖出的利润，并更新最大利润
		}
	}

	return maxProfit
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/
