/*
 * @lc app=leetcode.cn id=704 lang=golang
 * @lcpr version=30204
 *
 * [704] 二分查找
 */

// @lcpr-template-start
package leetcodespace

// @lcpr-template-end
// @lc code=start

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 初始化左右指针

	for left <= right { // 当左指针不越过右指针时继续查找
		mid := left + (right-left)/2 // 防止溢出的中间索引计算

		if nums[mid] == target {
			return mid // 找到目标值，返回其索引
		}

		if nums[mid] > target {
			right = mid - 1 // 目标值在左半部分，缩小右边界
		} else {
			left = mid + 1 // 目标值在右半部分，缩小左边界
		}
	}

	return -1 // 未找到目标值，返回 -1
}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,3,5,9,12]\n9\n
// @lcpr case=end

// @lcpr case=start
// [-1,0,3,5,9,12]\n2\n
// @lcpr case=end

*/
