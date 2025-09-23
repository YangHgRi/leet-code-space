/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30204
 *
 * [3] 无重复字符的最长子串
 */

// @lcpr-template-start
package leetcodespace

// @lcpr-template-end

// @lc code=start

func lengthOfLongestSubstring(s string) int {
	window := [128]bool{}
	ans := 0
	left := 0
	for right, c := range s {
		for window[c] {
			window[s[left]] = false
			left += 1
		}
		window[c] = true
		ans = max(ans, right-left+1)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
