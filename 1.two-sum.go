/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30204
 *
 * [1] 两数之和
 */

// @lcpr-template-start
package leet_code_space

// @lcpr-template-end
// @lc code=start

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)     // 创建一个哈希表 (map[int]int) 用于存储 (数字, 索引)
	for i, num := range nums { // 遍历切片，i 是索引，num 是值
		complement := target - num          // 计算补数
		if index, ok := m[complement]; ok { // 查找补数是否存在于哈希表中
			return []int{index, i} // 如果存在，返回当前索引和哈希表中存储的索引
		}
		m[num] = i // 将当前数字和它的索引存入哈希表
	}
	return []int{} // 如果没有找到，返回空切片
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/
