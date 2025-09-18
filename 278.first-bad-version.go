/*
 * @lc app=leetcode.cn id=278 lang=golang
 * @lcpr version=30204
 *
 * [278] 第一个错误的版本
 */

// @lcpr-template-start
package leet_code_space

import "math/rand"

var firstBadVersionNumber int = rand.Intn(10)

func isBadVersion(version int) bool {
	return version >= firstBadVersionNumber
}

// @lcpr-template-end
// @lc code=start

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	low := 1  // 搜索范围的最低版本号 (版本号从1开始)
	high := n // 搜索范围的最高版本号
	ans := n  // 记录找到的第一个错误版本，初始化为n（以防所有版本都错误）

	for low <= high {
		// 计算中间版本号，使用 (high-low)/2 + low 方式避免溢出
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			ans = mid      // mid 是一个错误版本，它可能是第一个，先记录下来
			high = mid - 1 // 尝试在 mid 之前寻找更早的错误版本
		} else {
			low = mid + 1 // mid 是一个好版本，所以第一个错误版本肯定在 mid 之后
		}
	}
	return ans // 循环结束后，ans 就是第一个错误版本
}

// @lc code=end

/*
// @lcpr case=start
// 5\n4\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
