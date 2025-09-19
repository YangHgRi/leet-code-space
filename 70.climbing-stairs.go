/*
 * @lc app=leetcode.cn id=70 lang=golang
 * @lcpr version=30204
 *
 * [70] 爬楼梯
 */

// @lcpr-template-start
package leetcodespace

// @lcpr-template-end
// @lc code=start

func climbStairs(n int) int {
	// 定义基本情况 (Base Cases)
	if n == 0 {
		return 0 // 通常此情况返回1 (不爬)，但这里遵循原Java代码逻辑
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// 初始化：last 对应 f(n-2) 的值, current 对应 f(n-1) 的值
	// 初始值分别为 f(2) 和 f(3)
	last := 2    // 爬2阶楼梯的方法数
	current := 3 // 爬3阶楼梯的方法数

	// 从 f(4) 开始迭代计算直到 f(n)
	// 循环次数为 n - 3 次
	for i := 0; i < n-3; i++ {
		temp := current // 保存 f(n-1) 的值
		current += last // 计算 f(n) = f(n-1) + f(n-2)
		last = temp     // 更新 f(n-2) 为之前的 f(n-1)
	}

	return current // 返回最终计算出的 f(n)
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/
