/*
 * @lc app=leetcode.cn id=9 lang=golang
 * @lcpr version=30204
 *
 * [9] 回文数
 */

// @lcpr-template-start
package leetcodespace

import "strconv"

// @lcpr-template-end
// @lc code=start

// 双指针遍历法
// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}
// 	if x == 0 {
// 		return true
// 	}
// 	// 优化：使用纯整数算术高效计算位数，避免math.Pow的潜在开销和浮点精度问题。
// 	place := 0 // `place` 将存储总位数
// 	temp := x
// 	for temp > 0 {
// 		temp /= 10
// 		place++
// 	}
// 	// lp: "左指针" 对应数字从右起的第1位 (个位, 10^0)
// 	// rp: "右指针" 对应数字从右起的第 `place` 位 (最高位, 10^(place-1))
// 	lp, rp := 1, place
// 	for lp < rp {
// 		// 获取当前左指针位置 (从右数第lp位) 的数字
// 		// 例如，对于12321, lp=1，digitL = (12321 / 10^0) % 10 = 1
// 		digitL := (x / int(math.Pow(10, float64(lp-1)))) % 10
// 		// 获取当前右指针位置 (从右数第rp位) 的数字，即从左数对应位置的数字
// 		// 例如，对于12321, rp=5，digitR = (12321 / 10^4) % 10 = 1
// 		digitR := (x / int(math.Pow(10, float64(rp-1)))) % 10
// 		if digitL != digitR {
// 			return false // 发现不匹配，立即返回
// 		}
// 		lp++ // 左指针向中间移动
// 		rp-- // 右指针向中间移动
// 	}
// 	return true // 所有检查的位都匹配，是回文数
// }

// 折叠对比法
// func isPalindrome(x int) bool {
// 	// 负数不是回文数
// 	// 以0结尾的非零数字，如10, 100，不可能是回文数，因为前导0是不允许的。
// 	// (0本身是回文数，已在下一个条件覆盖)
// 	if x < 0 || (x%10 == 0 && x != 0) {
// 		return false
// 	}

// 	revertedNumber := 0
// 	// 当原数字x大于反转数字revertedNumber时，持续反转
// 	// 当x小于或等于revertedNumber时，表示我们已经处理了一半或超过一半的数字
// 	for x > revertedNumber {
// 		revertedNumber = revertedNumber*10 + x%10 // 取末位，并按位构建反转数字
// 		x /= 10                                   // 移除原数字的末位
// 	}

// 	// 比较：
// 	// 1. 如果数字总位数为偶数（如1221），循环结束时 x 会等于 revertedNumber (12 == 12)
// 	// 2. 如果数字总位数为奇数（如12321），循环结束时 x 会等于 revertedNumber / 10 (12 == 123 / 10)
// 	//    因为中间位（3）在revertedNumber中，但我们只需比较两侧，所以丢弃中间位。
// 	return x == revertedNumber || x == revertedNumber/10
// }

// 字符串化遍历法,难度最低,时间复杂度反而容易刷满分,难绷
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// 121\n
// @lcpr case=end

// @lcpr case=start
// -121\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

*/
