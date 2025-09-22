/*
 * @lc app=leetcode.cn id=3005 lang=golang
 * @lcpr version=30204
 *
 * [3005] 最大频率元素计数
 */

// @lcpr-template-start
package leetcodespace

// @lcpr-template-end
// @lc code=start

func maxFrequencyElements(nums []int) int {
	freqMap := make(map[int]int)
	maxFreq := 0
	res := 0
	for _, num := range nums {
		freqMap[num]++
		currentNumFreq := freqMap[num]
		if currentNumFreq > maxFreq {
			maxFreq = currentNumFreq
			res = maxFreq
		} else if currentNumFreq == maxFreq {
			res += maxFreq
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,3,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

*/
