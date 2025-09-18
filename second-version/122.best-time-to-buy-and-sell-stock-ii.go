/*
 * @lc app=leetcode.cn id=122 lang=golang
 * @lcpr version=30204
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lcpr-template-start
package second_version

// @lcpr-template-end
// @lc code=start

func maxProfit(prices []int) int {
	maxProfit := 0

	// 遍历价格列表，直到倒数第二个元素
	// 因为需要比较 prices[i] 和 prices[i+1]
	for i := 0; i < len(prices)-1; i++ {
		// 如果今天的价格高于昨天的价格，就进行一次交易（买入昨天，卖出今天）
		// 并将利润累加
		if prices[i] < prices[i+1] {
			maxProfit += prices[i+1] - prices[i]
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
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/
